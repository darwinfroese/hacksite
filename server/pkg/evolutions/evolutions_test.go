package evolutions

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/darwinfroese/hacksite/server/models"
	"github.com/darwinfroese/hacksite/server/pkg/database"
	"github.com/darwinfroese/hacksite/server/pkg/database/bolt"
	"github.com/darwinfroese/hacksite/server/pkg/log"
	"github.com/darwinfroese/hacksite/server/pkg/log/testLogger"
)

var db database.Database
var logger log.Logger

var testProject = models.Project{
	ID:               "1234",
	Name:             "Test Project",
	CurrentEvolution: models.Evolution{},
}

func TestMain(m *testing.M) {
	db = bolt.New()
	logger = testLogger.New()
	err := db.AddProject(testProject)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't add the project to setup the tests: %s\n", err.Error())
		os.Exit(-1)
	}

	retCode := m.Run()

	os.Remove("database.db")
	os.Exit(retCode)
}

var createEvolutionTests = []struct {
	Description            string
	NewEvolution           models.Evolution
	ExpectedEvolutionCount int
}{{
	Description: "Creating a new evolution should add it to the list and set it as the next evolution",
	NewEvolution: models.Evolution{
		Number:    1,
		ProjectID: "1234",
		Tasks: []models.Task{
			models.Task{Task: "Test Task", ProjectID: "1234", Completed: false, EvolutionNumber: 1},
		},
	},
	ExpectedEvolutionCount: 1,
}, {
	Description: "Adding a second evolution to the project should replace the curent evolution and add it the list.",
	NewEvolution: models.Evolution{
		Number:    2,
		ProjectID: "1234",
		Tasks: []models.Task{
			models.Task{Task: "Second Evolution Task", ProjectID: "1234", Completed: true, EvolutionNumber: 2},
		},
	},
	ExpectedEvolutionCount: 2,
}}

func TestCreateEvolution(t *testing.T) {
	t.Log("Testing CreateEvolution...")

	for i, tc := range createEvolutionTests {
		t.Logf("[ %02d ] %s\n", i+1, tc.Description)

		project, err := CreateEvolution(db, logger, tc.NewEvolution)
		if err != nil {
			t.Errorf("[ FAIL ] Couldn't create the evolution: %s\n", err.Error())
		}

		if !reflect.DeepEqual(project.CurrentEvolution, tc.NewEvolution) {
			t.Error("[ FAIL ] The new evolution was not set as the current evolution.")
		}
		evolutionsCount := len(project.Evolutions)
		if evolutionsCount != tc.ExpectedEvolutionCount {
			t.Errorf("[ FAIL ] Received an unexpected number of evolutions.\nExpected: %d\nBut Got:  %d\n",
				tc.ExpectedEvolutionCount, evolutionsCount)
		}

		// update for persistance
		testProject = project
	}
}

var swapEvolutionsTests = []struct {
	Description   string
	SwapEvolution models.Evolution
	ShouldBeSet   bool
}{{
	Description: "Testing swapping an evolution should set the current evolution to the one selected.",
	SwapEvolution: models.Evolution{
		Number:    1,
		ProjectID: "1234",
		Tasks: []models.Task{
			models.Task{Task: "Test Task", Completed: false, ProjectID: "1234", ID: 0, EvolutionNumber: 1},
		},
	},
	ShouldBeSet: true,
}, {
	Description: "Attempting to swap to an evolution that does not exist should fail.",
	SwapEvolution: models.Evolution{
		Number:    10,
		ProjectID: "1234",
		Tasks: []models.Task{
			models.Task{Task: "Unknown evolution", Completed: true, ProjectID: "1234", ID: 0, EvolutionNumber: 10},
		},
	},
	ShouldBeSet: false,
}}

func TestSwapCurrentEvolution(t *testing.T) {
	t.Log("Testing SwapCurrentEvolution...")

	for i, tc := range swapEvolutionsTests {
		t.Logf("[ %02d ] %s\n", i+1, tc.Description)

		project, err := SwapCurrentEvolution(db, logger, tc.SwapEvolution)
		if err != nil && err.Error() != invalidEvolutionErrorMessage {
			t.Errorf("[ FAIL ] Couldn't swap the current evolution: %s\n", err.Error())
		}

		if reflect.DeepEqual(project.CurrentEvolution, tc.SwapEvolution) != tc.ShouldBeSet {
			t.Errorf("[ FAIL ] The current evolution was incorrectly updated.\nExpected: %+v\nBut Got:  %+v\n",
				project.CurrentEvolution, tc.SwapEvolution)
		}
	}
}
