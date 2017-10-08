package tasks

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/darwinfroese/hacksite/server/models"
	"github.com/darwinfroese/hacksite/server/pkg/database"
)

const testProjectID = 0

var db database.Database

var testProject = models.Project{
	Name: "TestProject",
	CurrentIteration: models.Iteration{
		Tasks: []models.Task{
			testTask,
		},
	},
}

var testTask = models.Task{
	Task:      "Test Task",
	Completed: false,
}

func TestMain(m *testing.M) {
	setup()

	retCode := m.Run()

	cleanup(retCode)
}

func setup() {
	db = database.CreateDB()

	err := db.AddProject(testProject)
	if err != nil {
		fmt.Printf("There was an error setting up the tests: %s\n", err.Error())
		cleanup(-1)
	}
}

func cleanup(exitCode int) {
	os.Remove("database.db")
	os.Exit(exitCode)
}

var updateTaskTests = []struct {
	Description          string
	Task                 models.Task
	ShouldContainNewTask bool
}{{
	Description: "Updating a task should replace the old task in it's project with the new one.",
	Task: models.Task{
		Task:      "This is an updated test task.",
		Completed: false,
	},
	ShouldContainNewTask: true,
}, {
	Description: "Updating a task's completion status should only change the completed status of the one task.",
	Task: models.Task{
		ID:        0,
		ProjectID: 0,
		Task:      "This is an updated test task.",
		Completed: true,
	},
	ShouldContainNewTask: true,
}, {
	Description: "Updating a task that does not exist in it's project should not update any tasks.",
	Task: models.Task{
		ID:        1234,
		ProjectID: 0,
		Task:      "This task doesn't exist",
		Completed: true,
	},
	ShouldContainNewTask: false,
}}

func TestUpdateTask(t *testing.T) {
	t.Log("Testing UpdateTask...")

	for i, tc := range updateTaskTests {
		t.Logf("[ %02d ] %s\n", i+1, tc.Description)

		project, err := UpdateTask(db, tc.Task)
		if err != nil {
			t.Errorf("[ FAIL ] Failed to update task: %s\n", err.Error())
		}

		if tc.ShouldContainNewTask && !contains(tc.Task, project.CurrentIteration.Tasks) {
			t.Error("[ FAIL ] Project should contain the new task but does not.")
		} else if !tc.ShouldContainNewTask && contains(tc.Task, project.CurrentIteration.Tasks) {
			t.Error("[ FAIL ] Project contained a task when it should not have contained the task.")
		}
	}
}

var removeTaskTests = []struct {
	Description       string
	TaskToDelete      models.Task
	ExpectedTask      models.Task
	ShouldContainTask bool
}{{
	Description: "Attempting to remove a task that doesn't exist in a project should leave existing tasks in the project.",
	TaskToDelete: models.Task{
		ID:        1234,
		ProjectID: 0,
	},
	ExpectedTask: models.Task{
		ID:        0,
		ProjectID: 0,
		Task:      "This is an updated test task.",
		Completed: true,
	},
	ShouldContainTask: true,
}, {
	Description: "After removing a task from a project the task should not be in the project.",
	TaskToDelete: models.Task{
		ID:        0,
		ProjectID: 0,
		Task:      "This is an updated test task.",
		Completed: true,
	},
	ShouldContainTask: false,
}}

func TestRemoveTask(t *testing.T) {
	t.Log("Testing RemoveTask...")

	for i, tc := range removeTaskTests {
		t.Logf("[ %02d ] %s\n", i+1, tc.Description)

		project, err := RemoveTask(db, tc.TaskToDelete)
		if err != nil {
			t.Errorf("[ FAIL ] Could not remove the task from the project: %s\n", err.Error())
		}

		if tc.ShouldContainTask && !contains(tc.ExpectedTask, project.CurrentIteration.Tasks) {
			t.Error("[ FAIL ] The wrong task was removed from the project.")
		} else if !tc.ShouldContainTask && contains(tc.ExpectedTask, project.CurrentIteration.Tasks) {
			t.Error("[ FAIL ] The task wasn't deleted from the database.")
		}
	}
}

func contains(t models.Task, list []models.Task) bool {
	for _, el := range list {
		if reflect.DeepEqual(t, el) {
			return true
		}
	}

	return false
}
