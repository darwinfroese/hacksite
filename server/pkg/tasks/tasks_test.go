package tasks

import (
	"os"
	"reflect"
	"testing"

	"github.com/darwinfroese/hacksite/server/models"
	"github.com/darwinfroese/hacksite/server/pkg/database"
)

var db database.Database
var testProject = models.Project{
	Name: "Test Project",
	CurrentIteration: models.Iteration{
		Number: 1,
		Tasks: []models.Task{
			models.Task{ID: 0, Task: "Test Task"},
			models.Task{ID: 1, Task: "Test Task 2"},
		},
	},
}

func TestMain(m *testing.M) {
	db = database.CreateDB()

	retCode := m.Run()

	os.Remove("database.db")
	os.Exit(retCode)
}

var updateTaskTests = []struct {
	Description     string
	NewTask         models.Task
	ExpectedProject models.Project
}{{
	Description: "Testing updating a task in a project returns the updated task in the project.",
	NewTask: models.Task{
		ID:        0,
		ProjectID: 0,
		Task:      "New Task",
		Completed: true,
	},
	ExpectedProject: models.Project{
		ID:     0,
		Name:   "Test Project",
		Status: models.StatusInProgress,
		CurrentIteration: models.Iteration{
			Number: 1,
			Tasks: []models.Task{
				models.Task{ID: 0, ProjectID: 0, Task: "New Task", Completed: true},
				models.Task{ID: 1, Task: "Test Task 2"},
			},
		},
	},
}, {
	Description: "Testing updating a task not in a project returns the existing project.",
	NewTask: models.Task{
		ID:        10,
		ProjectID: 0,
		Task:      "Non-Existent task",
		Completed: true,
	},
	ExpectedProject: models.Project{
		ID:     0,
		Name:   "Test Project",
		Status: models.StatusInProgress,
		CurrentIteration: models.Iteration{
			Number: 1,
			Tasks: []models.Task{
				models.Task{ID: 0, ProjectID: 0, Task: "New Task", Completed: true},
				models.Task{ID: 1, Task: "Test Task 2"},
			},
		},
	},
}}

func TestUpdateTask(t *testing.T) {
	t.Log("Testing UpdateTask...")
	db.AddProject(testProject)

	for i, tc := range updateTaskTests {
		t.Logf("[ %02d ] %s\n", i+1, tc.Description)

		proj, err := UpdateTask(db, tc.NewTask)
		if err != nil {
			t.Errorf("[ FAIL ] An error occured updating the task: %s\n", err.Error())
		}

		if !reflect.DeepEqual(proj, tc.ExpectedProject) {
			t.Errorf("[ FAIL ] The projects were not equal.\nExpected: %+v\nBut got:  %+v\n", tc.ExpectedProject, proj)
		}
	}

	db.RemoveProject(testProject.ID)
}

var removeTaskTests = []struct {
	Description     string
	TaskToRemove    models.Task
	ExpectedProject models.Project
}{{
	Description: "Testing removing a task from a project returns a project without the task.",
	TaskToRemove: models.Task{
		ID:        0,
		ProjectID: 0,
		Task:      "New Task",
		Completed: true,
	},
	ExpectedProject: models.Project{
		ID:     0,
		Name:   "Test Project",
		Status: models.StatusNew,
		CurrentIteration: models.Iteration{
			Number: 1,
			Tasks:  []models.Task{models.Task{ID: 1, Task: "Test Task 2"}},
		},
	},
}, {
	Description: "Testing removing a task from a project that doesn't exist returns the same project.",
	TaskToRemove: models.Task{
		ID:        10,
		ProjectID: 0,
		Task:      "Non Existant Task",
		Completed: true,
	},
	ExpectedProject: models.Project{
		ID:     0,
		Name:   "Test Project",
		Status: models.StatusNew,
		CurrentIteration: models.Iteration{
			Number: 1,
			Tasks:  []models.Task{models.Task{ID: 1, Task: "Test Task 2"}},
		},
	},
}}

func TestRemoveTask(t *testing.T) {
	t.Log("Testing RemoveTask...")
	db.AddProject(testProject)

	for i, tc := range removeTaskTests {
		t.Logf("[ %02d ] %s\n", i+1, tc.Description)

		proj, err := RemoveTask(db, tc.TaskToRemove)
		if err != nil {
			t.Errorf("[ FAIL ] An error occured updating the task: %s\n", err.Error())
		}

		if !reflect.DeepEqual(proj, tc.ExpectedProject) {
			t.Errorf("[ FAIL ] The projects were not equal.\nExpected: %+v\nBut got:  %+v\n", tc.ExpectedProject, proj)
		}
	}

	db.RemoveProject(testProject.ID)
}
