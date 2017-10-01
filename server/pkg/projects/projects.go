package projects

import (
	"fmt"
	"os"

	"github.com/darwinfroese/hacksite/server/models"
	"github.com/darwinfroese/hacksite/server/pkg/auth"
	"github.com/darwinfroese/hacksite/server/pkg/database"
)

// UpdateProjectStatus returns the string representation of the project's status
func UpdateProjectStatus(project models.Project) string {
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

// CreateProject grabs the next sequence in the database, sets up the project
// and inserts it into the database
func CreateProject(db database.Database, project *models.Project, session models.Session) error {
	// This is actually just setting the project status
	project.Status = UpdateProjectStatus(*project)

	id, err := db.GetNextProjectID()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		return err
	}

	project.ID = id
	project.CurrentIteration.ProjectID = project.ID
	project.CurrentIteration.Number = 1
	project.Iterations = append(project.Iterations, project.CurrentIteration)
	for i, task := range project.CurrentIteration.Tasks {
		task.ProjectID = project.ID
		task.IterationNumber = project.CurrentIteration.Number
		project.CurrentIteration.Tasks[i] = task
	}

	err = db.AddProject(*project)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		return err
	}

	return nil
}

// AddProjectToUser will add the project ID to the current users list
func AddProjectToUser(db database.Database, session models.Session, projectID int) error {
	account, err := auth.GetCurrentUser(db, session)
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
