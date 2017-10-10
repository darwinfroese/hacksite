package iterations

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/darwinfroese/hacksite/server/models"
	"github.com/darwinfroese/hacksite/server/pkg/database"
	"github.com/darwinfroese/hacksite/server/pkg/database/bolt"
)

var db database.Database

var testProject = models.Project{
	Name:             "Test Project",
	CurrentIteration: models.Iteration{},
}

func TestMain(m *testing.M) {
	db = bolt.New()
	err := db.AddProject(testProject)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't add the project to setup the tests: %s\n", err.Error())
		os.Exit(-1)
	}

	retCode := m.Run()

	os.Remove("database.db")
	os.Exit(retCode)
}

var createIterationTests = []struct {
	Description            string
	NewIteration           models.Iteration
	ExpectedIterationCount int
}{{
	Description: "Creating a new iteration should add it to the list and set it as the next iteration",
	NewIteration: models.Iteration{
		Number:    1,
		ProjectID: 0,
		Tasks: []models.Task{
			models.Task{Task: "Test Task", Completed: false, IterationNumber: 1},
		},
	},
	ExpectedIterationCount: 1,
}, {
	Description: "Adding a second iteration to the project should replace the curent iteration and add it the list.",
	NewIteration: models.Iteration{
		Number:    2,
		ProjectID: 0,
		Tasks: []models.Task{
			models.Task{Task: "Second Iteration Task", Completed: true, IterationNumber: 2},
		},
	},
	ExpectedIterationCount: 2,
}}

func TestCreateIteration(t *testing.T) {
	t.Log("Testing CreateIteration...")

	for i, tc := range createIterationTests {
		t.Logf("[ %02d ] %s\n", i+1, tc.Description)

		project, err := CreateIteration(db, tc.NewIteration)
		if err != nil {
			t.Errorf("[ FAIL ] Couldn't create the iteration: %s\n", err.Error())
		}

		if !reflect.DeepEqual(project.CurrentIteration, tc.NewIteration) {
			t.Error("[ FAIL ] The new iteration was not set as the current iteration.")
		}
		iterationsCount := len(project.Iterations)
		if iterationsCount != tc.ExpectedIterationCount {
			t.Errorf("[ FAIL ] Received an unexpected number of iterations.\nExpected: %d\nBut Got:  %d\n",
				tc.ExpectedIterationCount, iterationsCount)
		}

		// update for persistance
		testProject = project
	}
}

var swapIterationsTests = []struct {
	Description   string
	SwapIteration models.Iteration
	ShouldBeSet   bool
}{{
	Description: "Testing swapping an iteration should set the current iteration to the one selected.",
	SwapIteration: models.Iteration{
		Number:    1,
		ProjectID: 0,
		Tasks: []models.Task{
			models.Task{Task: "Test Task", Completed: false, ProjectID: 0, ID: 0, IterationNumber: 1},
		},
	},
	ShouldBeSet: true,
}, {
	Description: "Attempting to swap to an iteration that does not exist should fail.",
	SwapIteration: models.Iteration{
		Number:    10,
		ProjectID: 0,
		Tasks: []models.Task{
			models.Task{Task: "Unknown iteration", Completed: true, ProjectID: 0, ID: 0, IterationNumber: 10},
		},
	},
	ShouldBeSet: false,
}}

func TestSwapCurrentIteration(t *testing.T) {
	t.Log("Testing SwapCurrentIteration...")

	for i, tc := range swapIterationsTests {
		t.Logf("[ %02d ] %s\n", i+1, tc.Description)

		project, err := SwapCurrentIteration(db, tc.SwapIteration)
		if err != nil && err.Error() != models.InvalidIterationErrorMessage {
			t.Errorf("[ FAIL ] Couldn't swap the current iteration: %s\n", err.Error())
		}

		if reflect.DeepEqual(project.CurrentIteration, tc.SwapIteration) != tc.ShouldBeSet {
			t.Error("[ FAIL ] The current iteration was incorrectly updated.")
		}
	}
}
