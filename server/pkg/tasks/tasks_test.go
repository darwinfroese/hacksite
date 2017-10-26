package tasks

import (
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
	ID:   "1234",
	Name: "Test Project",
	CurrentEvolution: models.Evolution{
		Number:    1,
		ProjectID: "1234",
		Tasks: []models.Task{
			models.Task{ID: 0, ProjectID: "1234", Task: "Test Task"},
			models.Task{ID: 1, ProjectID: "1234", Task: "Test Task 2"},
		},
	},
}

var multiEvolutionTestProject = models.Project{
	ID:   "4321",
	Name: "Test Project 2",
	CurrentEvolution: models.Evolution{
		Number:    1,
		ProjectID: "4321",
		Tasks: []models.Task{
			models.Task{ID: 0, ProjectID: "4321", Task: "Test Task", Completed: true},
			models.Task{ID: 1, ProjectID: "4321", Task: "Test Task 2"},
		},
	},
	Evolutions: []models.Evolution{
		models.Evolution{
			Number:    1,
			ProjectID: "4321",
			Tasks: []models.Task{
				models.Task{ID: 0, ProjectID: "4321", Task: "Test Task", Completed: true},
				models.Task{ID: 1, ProjectID: "4321", Task: "Test Task 2"},
			},
		},
		models.Evolution{
			Number:    2,
			ProjectID: "4321",
			Tasks: []models.Task{
				models.Task{ID: 0, ProjectID: "4321", Task: "Test Task 5", Completed: true},
				models.Task{ID: 1, ProjectID: "4321", Task: "Test Task 6", Completed: true},
			},
		},
		models.Evolution{
			Number:    3,
			ProjectID: "4321",
			Tasks: []models.Task{
				models.Task{ID: 0, ProjectID: "4321", Task: "Test Task 3"},
				models.Task{ID: 1, ProjectID: "4321", Task: "Test Task 4"},
			},
		},
	},
}

func TestMain(m *testing.M) {
	db = bolt.New()
	logger = testLogger.New()

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
		ProjectID: "1234",
		Task:      "New Task",
		Completed: true,
	},
	ExpectedProject: models.Project{
		ID:     "1234",
		Name:   "Test Project",
		Status: "InProgress",
		CurrentEvolution: models.Evolution{
			ProjectID: "1234",
			Number:    1,
			Tasks: []models.Task{
				models.Task{ID: 0, ProjectID: "1234", Task: "New Task", Completed: true},
				models.Task{ID: 1, ProjectID: "1234", Task: "Test Task 2"},
			},
		},
	},
}, {
	Description: "Testing updating a task not in a project returns the existing project.",
	NewTask: models.Task{
		ID:        10,
		ProjectID: "1234",
		Task:      "Non-Existent task",
		Completed: true,
	},
	ExpectedProject: models.Project{
		ID:     "1234",
		Name:   "Test Project",
		Status: "InProgress",
		CurrentEvolution: models.Evolution{
			Number:    1,
			ProjectID: "1234",
			Tasks: []models.Task{
				models.Task{ID: 0, ProjectID: "1234", Task: "New Task", Completed: true},
				models.Task{ID: 1, ProjectID: "1234", Task: "Test Task 2"},
			},
		},
	},
}, {
	Description: "Testing completing the last task should mark the project as completed.",
	NewTask: models.Task{
		ID:        1,
		ProjectID: "1234",
		Task:      "New Task 2",
		Completed: true,
	},
	ExpectedProject: models.Project{
		ID:     "1234",
		Name:   "Test Project",
		Status: "Completed",
		CurrentEvolution: models.Evolution{
			ProjectID: "1234",
			Number:    1,
			Tasks: []models.Task{
				models.Task{ID: 0, ProjectID: "1234", Task: "New Task", Completed: true},
				models.Task{ID: 1, ProjectID: "1234", Task: "New Task 2", Completed: true},
			},
		},
	},
}, {
	Description: "Testing completing the last task should switch to the next evolution if it exists.",
	NewTask: models.Task{
		ID:        1,
		ProjectID: "4321",
		Task:      "Test Task 2",
		Completed: true,
	},
	ExpectedProject: models.Project{
		ID:     "4321",
		Name:   "Test Project 2",
		Status: "InProgress",
		CurrentEvolution: models.Evolution{
			Number:    3,
			ProjectID: "4321",
			Tasks: []models.Task{
				models.Task{ID: 0, ProjectID: "4321", Task: "Test Task 3"},
				models.Task{ID: 1, ProjectID: "4321", Task: "Test Task 4"},
			},
		},
		Evolutions: []models.Evolution{
			models.Evolution{
				Number:    1,
				ProjectID: "4321",
				Tasks: []models.Task{
					models.Task{ID: 0, ProjectID: "4321", Task: "Test Task", Completed: true},
					models.Task{ID: 1, ProjectID: "4321", Task: "Test Task 2", Completed: true},
				},
			},
			models.Evolution{
				Number:    2,
				ProjectID: "4321",
				Tasks: []models.Task{
					models.Task{ID: 0, ProjectID: "4321", Task: "Test Task 5", Completed: true},
					models.Task{ID: 1, ProjectID: "4321", Task: "Test Task 6", Completed: true},
				},
			},
			models.Evolution{
				Number:    3,
				ProjectID: "4321",
				Tasks: []models.Task{
					models.Task{ID: 0, ProjectID: "4321", Task: "Test Task 3"},
					models.Task{ID: 1, ProjectID: "4321", Task: "Test Task 4"},
				},
			},
		},
	},
}}

func TestUpdateTask(t *testing.T) {
	t.Log("Testing UpdateTask...")
	db.AddProject(testProject)
	db.AddProject(multiEvolutionTestProject)

	for i, tc := range updateTaskTests {
		t.Logf("[ %02d ] %s\n", i+1, tc.Description)

		proj, err := UpdateTask(db, logger, tc.NewTask)
		if err != nil {
			t.Errorf("[ FAIL ] An error occured updating the task: %s\n", err.Error())
		}

		if !reflect.DeepEqual(proj, tc.ExpectedProject) {
			t.Errorf("[ FAIL ] The projects were not equal.\nExpected: %+v\nBut got:  %+v\n", tc.ExpectedProject, proj)
		}
	}

	db.RemoveProject(testProject.ID)
	db.RemoveProject(multiEvolutionTestProject.ID)
}

var removeTaskTests = []struct {
	Description     string
	TaskToRemove    models.Task
	ExpectedProject models.Project
}{{
	Description: "Testing removing a task from a project returns a project without the task.",
	TaskToRemove: models.Task{
		ID:        0,
		ProjectID: "1234",
		Task:      "New Task",
		Completed: true,
	},
	ExpectedProject: models.Project{
		ID:     "1234",
		Name:   "Test Project",
		Status: "New",
		CurrentEvolution: models.Evolution{
			Number:    1,
			ProjectID: "1234",
			Tasks:     []models.Task{models.Task{ID: 1, ProjectID: "1234", Task: "Test Task 2"}},
		},
	},
}, {
	Description: "Testing removing a task from a project that doesn't exist returns the same project.",
	TaskToRemove: models.Task{
		ID:        10,
		ProjectID: "1234",
		Task:      "Non Existant Task",
		Completed: true,
	},
	ExpectedProject: models.Project{
		ID:     "1234",
		Name:   "Test Project",
		Status: "New",
		CurrentEvolution: models.Evolution{
			ProjectID: "1234",
			Number:    1,
			Tasks:     []models.Task{models.Task{ID: 1, ProjectID: "1234", Task: "Test Task 2"}},
		},
	},
}}

func TestRemoveTask(t *testing.T) {
	t.Log("Testing RemoveTask...")
	db.AddProject(testProject)

	for i, tc := range removeTaskTests {
		t.Logf("[ %02d ] %s\n", i+1, tc.Description)

		proj, err := RemoveTask(db, logger, tc.TaskToRemove)
		if err != nil {
			t.Errorf("[ FAIL ] An error occured updating the task: %s\n", err.Error())
		}

		if !reflect.DeepEqual(proj, tc.ExpectedProject) {
			t.Errorf("[ FAIL ] The projects were not equal.\nExpected: %+v\nBut got:  %+v\n", tc.ExpectedProject, proj)
		}
	}

	db.RemoveProject(testProject.ID)
}
