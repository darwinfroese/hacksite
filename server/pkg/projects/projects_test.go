package projects

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
var testSession = models.Session{
	Token:    "TestSession",
	Username: "test-account",
}

var projectID string

func TestMain(m *testing.M) {
	db = bolt.New()
	logger = testLogger.New()
	db.CreateAccount(models.Account{
		Username: "test-account",
	})

	retCode := m.Run()

	os.Remove("database.db")
	os.Exit(retCode)
}

var createProjectTests = []struct {
	Description     string
	ProjectToCreate models.Project
	ExpectedProject models.Project
}{{
	Description: "Testing that creating a project with a valid model returns a new project.",
	ProjectToCreate: models.Project{
		Name:             "test-project",
		CurrentEvolution: models.Evolution{},
	},
	ExpectedProject: models.Project{
		Name: "test-project",
		CurrentEvolution: models.Evolution{
			Number: 1,
		},
		Evolutions: []models.Evolution{
			models.Evolution{Number: 1},
		},
		Status: statusCompleted,
	},
}}

func TestCreateProject(t *testing.T) {
	t.Log("Testing CreateProject...")

	for i, tc := range createProjectTests {
		t.Logf("[ %02d ] %s\n", i+1, tc.Description)

		err := CreateProject(db, logger, &tc.ProjectToCreate, testSession.Username)
		if err != nil {
			t.Errorf("[ FAIL ] An unexpected error occured creating the project: %s\n", err.Error())
		}

		// ID is randomly generated so we need to copy it
		tc.ExpectedProject.ID = tc.ProjectToCreate.ID
		tc.ExpectedProject.CurrentEvolution.ProjectID = tc.ProjectToCreate.ID
		tc.ExpectedProject.Evolutions[0].ProjectID = tc.ProjectToCreate.ID

		if !reflect.DeepEqual(tc.ProjectToCreate, tc.ExpectedProject) {
			t.Errorf("[ FAIL ] The project created has unexpected values.\nExpected: %+v\nBut got:  %+v\n",
				tc.ExpectedProject, tc.ProjectToCreate)
		}
	}
}

var getUserProjectsTests = []struct {
	Description          string
	ExpectedProjectCount int
}{{
	Description:          "Testing that only the correct projects are returned for the user.",
	ExpectedProjectCount: 1,
}}

func TestGetUserProjects(t *testing.T) {
	t.Log("Testing GetUserProjects")

	for i, tc := range getUserProjectsTests {
		t.Logf("[ %02d ] %s\n", i+1, tc.Description)

		projects, err := GetUserProjects(db, logger, testSession.Username)
		if err != nil {
			t.Errorf("[ FAIL ] An unexpected error getting the projects for a user: %s\n", err.Error())
		}

		projectCount := len(projects)
		if projectCount != tc.ExpectedProjectCount {
			t.Errorf("[ FAIL ] Did not receive the number of projects expected. Expected: %d but got %d.\n", tc.ExpectedProjectCount, projectCount)
		}
	}
}

var updateProjectTests = []struct {
	Description     string
	NewProject      models.Project
	ExpectedProject models.Project
}{{
	Description: "Testing updating a project returns the new project as expected.",
	NewProject: models.Project{
		ID:   "1",
		Name: "test-project",
		CurrentEvolution: models.Evolution{
			Number: 1, ProjectID: "1",
			Tasks: []models.Task{
				models.Task{Task: "test", EvolutionNumber: 1},
			},
		},
		Evolutions: []models.Evolution{
			models.Evolution{
				Number: 1, ProjectID: "1",
				Tasks: []models.Task{
					models.Task{Task: "test", EvolutionNumber: 1},
				},
			},
		},
	},
	ExpectedProject: models.Project{
		ID:   "1",
		Name: "test-project",
		CurrentEvolution: models.Evolution{
			Number: 1, ProjectID: "1",
			Tasks: []models.Task{
				models.Task{Task: "test", EvolutionNumber: 1},
			},
		},
		Evolutions: []models.Evolution{
			models.Evolution{
				Number: 1, ProjectID: "1",
				Tasks: []models.Task{
					models.Task{Task: "test", EvolutionNumber: 1},
				},
			},
		},
		Status: statusNew,
	},
}}

func TestUpdateProject(t *testing.T) {
	t.Log("Testing UpdateProject...")

	for i, tc := range updateProjectTests {
		t.Logf("[ %02d ] %s\n", i+1, tc.Description)

		err := UpdateProject(db, &tc.NewProject)
		if err != nil {
			t.Errorf("[ FAIL ] An unexpected error occurred updating the project: %s\n", err.Error())
		}

		if !reflect.DeepEqual(tc.NewProject, tc.ExpectedProject) {
			t.Errorf("[ FAIL ] UpdateProject did not return the expected project.\nExpected: %+v\nBut got:  %+v\n", tc.ExpectedProject, tc.NewProject)
		}
	}
}

var deleteProjectTests = []struct {
	Description   string
	ExpectedCount int
}{{
	Description:   "Testing removing a project removes project from DB and user list.",
	ExpectedCount: 0,
}}

func TestDeleteProject(t *testing.T) {
	t.Log("Testing DeleteProject...")

	for i, tc := range deleteProjectTests {
		t.Logf("[ %02d ] %s\n", i+1, tc.Description)

		account, _ := db.GetAccountByUsername(testSession.Username)
		err := DeleteProject(db, logger, testSession.Username, account.ProjectIds[0])
		if err != nil {
			t.Errorf("[ FAIL ] An unexpected error occurred deleting the project: %s.\n", err.Error())
		}

		account, _ = db.GetAccountByUsername(testSession.Username)
		projectCount := len(account.ProjectIds)

		if projectCount != tc.ExpectedCount {
			t.Error("[ FAIL ] The project was not removed from the users list.\n")
		}
	}
}

var removeIDFromListTests = []struct {
	Description          string
	IDToRemove           string
	IDList, ExpectedList []string
}{{
	Description:  "Testing removing the first ID succeeds.",
	IDToRemove:   "1",
	IDList:       []string{"1", "2", "3", "4"},
	ExpectedList: []string{"2", "3", "4"},
}, {
	Description:  "Testing removing the last ID succeeds.",
	IDToRemove:   "4",
	IDList:       []string{"1", "2", "3", "4"},
	ExpectedList: []string{"1", "2", "3"},
}, {
	Description:  "Testing removing an ID in the middle succeeds.",
	IDToRemove:   "3",
	IDList:       []string{"1", "2", "3", "4"},
	ExpectedList: []string{"1", "2", "4"},
}, {
	Description:  "Testing attempting to remove an ID not in the list returns the same list.",
	IDToRemove:   "5",
	IDList:       []string{"1", "2", "3", "4"},
	ExpectedList: []string{"1", "2", "3", "4"},
}, {
	Description:  "Testing removing the only ID succeeds.",
	IDToRemove:   "1",
	IDList:       []string{"1"},
	ExpectedList: []string{},
}}

func TestRemoveIDFromList(t *testing.T) {
	t.Log("Testing removeIDFromList...")

	for i, tc := range removeIDFromListTests {
		t.Logf("[ %02d ] %s\n", i+1, tc.Description)

		newList := removeIDFromList(tc.IDToRemove, tc.IDList)
		if !reflect.DeepEqual(tc.ExpectedList, newList) {
			t.Errorf("[ FAIL ] Did not receive the list expected.\nExpected: %v\nBut got:  %v\n", tc.ExpectedList, newList)
		}
	}
}

//Validator TEST
var validateProjectTests = []struct {
	Description    string
	Project        models.Project
	ExpectedResult bool
}{{
	Description: "Testing a valid project model should validate.",
	Project: models.Project{
		Name: "Test Project",
	},
	ExpectedResult: true,
}, {
	Description:    "Testing a project missing a name value should not validate.",
	Project:        models.Project{},
	ExpectedResult: false,
}}

func TestValidateProject(t *testing.T) {
	t.Log("Testing ValidateProject...")

	for i, tc := range validateProjectTests {
		t.Logf("[ %02d ] %s\n", i+1, tc.Description)
		p := tc.Project
		err := p.Validate()
		resultValidateMethod := true
		result := tc.ExpectedResult

		if err != nil {
			resultValidateMethod = false
		}

		if result != resultValidateMethod {
			t.Errorf("[ FAIL ] ValidateProject did not return expected value. Expected \"%v\" but got \"%v\".",
				result, resultValidateMethod)
		}
	}
}
