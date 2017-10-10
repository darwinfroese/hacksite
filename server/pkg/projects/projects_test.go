package projects

import (
	"os"
	"reflect"
	"testing"

	"github.com/darwinfroese/hacksite/server/models"
	"github.com/darwinfroese/hacksite/server/pkg/database"
	"github.com/darwinfroese/hacksite/server/pkg/database/bolt"
)

var db database.Database
var testSession = models.Session{
	Token:  "TestSession",
	UserID: 1,
}

func TestMain(m *testing.M) {
	db = bolt.New()
	db.CreateAccount(models.Account{
		ID: 1,
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
		CurrentIteration: models.Iteration{},
	},
	ExpectedProject: models.Project{
		ID:   1,
		Name: "test-project",
		CurrentIteration: models.Iteration{
			Number: 1, ProjectID: 1,
		},
		Iterations: []models.Iteration{
			models.Iteration{Number: 1, ProjectID: 1},
		},
		Status: models.StatusCompleted,
	},
}}

func TestCreateProject(t *testing.T) {
	t.Log("Testing CreateProject...")

	for i, tc := range createProjectTests {
		t.Logf("[ %02d ] %s\n", i+1, tc.Description)

		err := CreateProject(db, &tc.ProjectToCreate, testSession.UserID)
		if err != nil {
			t.Errorf("[ FAIL ] An unexpected error occured creating the project: %s\n", err.Error())
		}

		if !reflect.DeepEqual(tc.ProjectToCreate, tc.ExpectedProject) {
			t.Errorf("[ FAIL ] The project created has unexpected values.\nExpected: %+v\nBut got:  %+v\n",
				tc.ExpectedProject, tc.ProjectToCreate)
		}
	}
}

var getUserProjectsTests = []struct {
	Description          string
	ExpectedProjectCount int
	ExpectedProjectIDs   []int
}{{
	Description:          "Testing that only the correct projects are returned for the user.",
	ExpectedProjectCount: 1,
	ExpectedProjectIDs:   []int{1},
}}

func TestGetUserProjects(t *testing.T) {
	t.Log("Testing GetUserProjects")

	for i, tc := range getUserProjectsTests {
		t.Logf("[ %02d ] %s\n", i+1, tc.Description)

		projects, err := GetUserProjects(db, testSession.UserID)
		if err != nil {
			t.Errorf("[ FAIL ] An unexpected error getting the projects for a user: %s\n", err.Error())
		}

		projectCount := len(projects)
		if projectCount != tc.ExpectedProjectCount {
			t.Errorf("[ FAIL ] Did not receive the number of projects expected. Expected: %d but got %d.\n", tc.ExpectedProjectCount, projectCount)
		}

		var projectIDs = []int{}
		for _, project := range projects {
			projectIDs = append(projectIDs, project.ID)
		}

		if !reflect.DeepEqual(projectIDs, tc.ExpectedProjectIDs) {
			t.Errorf("[ FAIL ] Did not receive the projects expected.\nExpected: %v\nBut got:  %v\n", tc.ExpectedProjectIDs, projectIDs)
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
		ID:   1,
		Name: "test-project",
		CurrentIteration: models.Iteration{
			Number: 1, ProjectID: 1,
			Tasks: []models.Task{
				models.Task{Task: "test", IterationNumber: 1},
			},
		},
		Iterations: []models.Iteration{
			models.Iteration{
				Number: 1, ProjectID: 1,
				Tasks: []models.Task{
					models.Task{Task: "test", IterationNumber: 1},
				},
			},
		},
	},
	ExpectedProject: models.Project{
		ID:   1,
		Name: "test-project",
		CurrentIteration: models.Iteration{
			Number: 1, ProjectID: 1,
			Tasks: []models.Task{
				models.Task{Task: "test", IterationNumber: 1},
			},
		},
		Iterations: []models.Iteration{
			models.Iteration{
				Number: 1, ProjectID: 1,
				Tasks: []models.Task{
					models.Task{Task: "test", IterationNumber: 1},
				},
			},
		},
		Status: models.StatusNew,
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
	Description       string
	ProjectIDToRemove int
	ExpectedCount     int
}{{
	Description:       "Testing removing a project removes project from DB and user list.",
	ProjectIDToRemove: 1,
	ExpectedCount:     0,
}}

func TestDeleteProject(t *testing.T) {
	t.Log("Testing DeleteProject...")

	for i, tc := range deleteProjectTests {
		t.Logf("[ %02d ] %s\n", i+1, tc.Description)

		err := DeleteProject(db, testSession.UserID, tc.ProjectIDToRemove)
		if err != nil {
			t.Errorf("[ FAIL ] An unexpected error occurred deleting the project: %s.\n", err.Error())
		}

		account, _ := db.GetAccountByID(1)
		projectCount := len(account.ProjectIds)

		if projectCount != tc.ExpectedCount {
			t.Error("[ FAIL ] The project was not removed from the users list.\n")
		}
	}
}

var removeIDFromListTests = []struct {
	Description          string
	IDToRemove           int
	IDList, ExpectedList []int
}{{
	Description:  "Testing removing the first ID succeeds.",
	IDToRemove:   1,
	IDList:       []int{1, 2, 3, 4},
	ExpectedList: []int{2, 3, 4},
}, {
	Description:  "Testing removing the last ID succeeds.",
	IDToRemove:   4,
	IDList:       []int{1, 2, 3, 4},
	ExpectedList: []int{1, 2, 3},
}, {
	Description:  "Testing removing an ID in the middle succeeds.",
	IDToRemove:   3,
	IDList:       []int{1, 2, 3, 4},
	ExpectedList: []int{1, 2, 4},
}, {
	Description:  "Testing attempting to remove an ID not in the list returns the same list.",
	IDToRemove:   5,
	IDList:       []int{1, 2, 3, 4},
	ExpectedList: []int{1, 2, 3, 4},
}, {
	Description:  "Testing removing the only ID succeeds.",
	IDToRemove:   1,
	IDList:       []int{1},
	ExpectedList: []int{},
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
