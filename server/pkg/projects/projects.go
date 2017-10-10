package projects

import (
	"fmt"
	"os"

	"github.com/darwinfroese/hacksite/server/models"
	"github.com/darwinfroese/hacksite/server/pkg/database"
)

// GetUserProjects grabs the project from database
func GetUserProjects(db database.Database, userID int) ([]models.Project, error) {
	account, err := db.GetAccountByID(userID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		return nil, err
	}

	var projects []models.Project
	for _, id := range account.ProjectIds {
		p, err := db.GetProject(id)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error Geting project from DB: %s\n", err.Error())
		} else {
			projects = append(projects, p)
		}
	}

	return projects, nil
}

// CreateProject grabs the next sequence in the database, sets up the project
// and inserts it into the database. CreateProject assumes the model has already
// been validated.
func CreateProject(db database.Database, project *models.Project, session models.Session) error {
	// This is actually just setting the project status
	project.Status = updateProjectStatus(*project)

	id, err := db.GetNextProjectID()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		return err
	}

	project.ID = id
	updateIteration(project.ID, 1, &project.CurrentIteration)
	project.Iterations = append(project.Iterations, project.CurrentIteration)
	project.CurrentIteration.Tasks = updateTasks(
		project.ID, project.CurrentIteration.Number, project.CurrentIteration.Tasks)

	addProjectToUser(db, session.UserID, id)
	err = db.AddProject(*project)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		return err
	}

	return nil
}

// DeleteProject will remove the project from the database as well as the users list of projects
func DeleteProject(db database.Database, userID, projectID int) error {
	err := db.RemoveProject(projectID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		return err
	}

	account, err := db.GetAccountByID(userID)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		return err
	}

	account.ProjectIds = removeIDFromList(projectID, account.ProjectIds)
	err = db.UpdateAccount(account)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		return err
	}

	return nil
}

// UpdateProject will update the status and make the change in the database as well
func UpdateProject(db database.Database, project *models.Project) error {
	project.Status = updateProjectStatus(*project)

	return db.UpdateProject(*project)
}

// HelperFunctions
func updateIteration(id, number int, iteration *models.Iteration) {
	iteration.ProjectID = id
	iteration.Number = number
}

func updateTasks(id, number int, tasks []models.Task) []models.Task {
	var newTasks []models.Task

	for _, t := range tasks {
		t.ProjectID = id
		t.IterationNumber = number

		newTasks = append(newTasks, t)
	}

	return newTasks
}

// addProjectToUser will add the project ID to the current users list
func addProjectToUser(db database.Database, userID, projectID int) error {
	account, err := db.GetAccountByID(userID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		return err
	}

	account.ProjectIds = append(account.ProjectIds, projectID)
	err = db.UpdateAccount(account)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		return err
	}

	return nil
}

func removeIDFromList(idToRemove int, idList []int) []int {
	for i, id := range idList {
		if id == idToRemove {
			list := append(idList[:i], idList[i+1:]...)
			return list
		}
	}

	return idList
}

func updateProjectStatus(project models.Project) string {
	complete := 0
	status := models.StatusNew

	tasks := project.CurrentIteration.Tasks
	for _, task := range tasks {
		if task.Completed {
			complete++
		}
	}
	if complete == len(tasks) {
		status = models.StatusCompleted
	} else if complete > 0 {
		status = models.StatusInProgress
	} else {
		status = models.StatusNew
	}

	return status
}
