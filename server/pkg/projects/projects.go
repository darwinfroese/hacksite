package projects

import (
	"fmt"

	"github.com/nu7hatch/gouuid"

	"github.com/darwinfroese/hacksite/server/models"
	"github.com/darwinfroese/hacksite/server/pkg/database"
	"github.com/darwinfroese/hacksite/server/pkg/log"
)

const (
	statusCompleted  = "Completed"
	statusInProgress = "InProgress"
	statusNew        = "New"
)

// GetUserProjects grabs the project from database
func GetUserProjects(db database.Database, logger log.Logger, username string) ([]models.Project, error) {
	account, err := db.GetAccountByUsername(username)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	var projects []models.Project
	for _, id := range account.ProjectIds {
		p, err := db.GetProject(id)
		if err != nil {
			logger.Error(fmt.Sprintf("Error Geting project from DB: %s\n", err.Error()))
		} else {
			projects = append(projects, p)
		}
	}

	return projects, nil
}

// CreateProject grabs the next sequence in the database, sets up the project
// and inserts it into the database. CreateProject assumes the model has already
// been validated.
func CreateProject(db database.Database, logger log.Logger, project *models.Project, username string) error {
	// This is actually just setting the project status
	project.Status = updateProjectStatus(project.CurrentEvolution, len(project.Evolutions))

	id, err := uuid.NewV4()
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	project.ID = id.String()
	updateEvolution(project.ID, 1, &project.CurrentEvolution)
	project.Evolutions = append(project.Evolutions, project.CurrentEvolution)
	project.CurrentEvolution.Tasks = updateTasks(
		project.ID, project.CurrentEvolution.Number, project.CurrentEvolution.Tasks)

	addProjectToUser(db, logger, username, id.String())
	err = db.AddProject(*project)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}

// DeleteProject will remove the project from the database as well as the users list of projects
func DeleteProject(db database.Database, logger log.Logger, username, projectID string) error {
	err := db.RemoveProject(projectID)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	account, err := db.GetAccountByUsername(username)

	if err != nil {
		logger.Error(err.Error())
		return err
	}

	account.ProjectIds = removeIDFromList(projectID, account.ProjectIds)
	err = db.UpdateAccount(account)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}

// UpdateProject will update the status and make the change in the database as well
func UpdateProject(db database.Database, project *models.Project) error {
	project.Status = updateProjectStatus(project.CurrentEvolution, len(project.Evolutions))
	swapEvolution(project)

	return db.UpdateProject(*project)
}

// HelperFunctions
func updateEvolution(id string, number int, evolution *models.Evolution) {
	evolution.ProjectID = id
	evolution.Number = number
}

func updateTasks(id string, number int, tasks []models.Task) []models.Task {
	var newTasks []models.Task

	for _, t := range tasks {
		t.ProjectID = id
		t.EvolutionNumber = number

		newTasks = append(newTasks, t)
	}

	return newTasks
}

// addProjectToUser will add the project ID to the current users list
func addProjectToUser(db database.Database, logger log.Logger, username, projectID string) error {
	account, err := db.GetAccountByUsername(username)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	account.ProjectIds = append(account.ProjectIds, projectID)
	err = db.UpdateAccount(account)

	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}

func removeIDFromList(idToRemove string, idList []string) []string {
	for i, id := range idList {
		if id == idToRemove {
			list := append(idList[:i], idList[i+1:]...)
			return list
		}
	}

	return idList
}

func updateProjectStatus(evolution models.Evolution, evolutionCount int) string {
	complete := 0
	status := statusNew

	tasks := evolution.Tasks
	for _, task := range tasks {
		if task.Completed {
			complete++
		}
	}
	if complete == len(tasks) {
		status = statusCompleted
	} else if complete > 0 {
		status = statusInProgress
	} else if complete == 0 && evolutionCount > 1 {
		status = statusInProgress
	} else {
		status = statusNew
	}

	return status
}

func swapEvolution(project *models.Project) {
	if project.Status != statusCompleted {
		return
	}

	// is this the last iteration
	for _, ev := range project.Evolutions {
		if ev.Number > project.CurrentEvolution.Number && updateProjectStatus(ev, 0) != statusCompleted {
			project.Status = statusInProgress
			project.CurrentEvolution = ev
			return
		}
	}
}
