package iterations

import (
	"errors"
	"reflect"

	"github.com/darwinfroese/hacksite/server/models"
	"github.com/darwinfroese/hacksite/server/pkg/database"
	"github.com/darwinfroese/hacksite/server/pkg/log"
	"github.com/darwinfroese/hacksite/server/pkg/projects"
)

// CreateIteration creates a new iteration and stores it in the database
func CreateIteration(db database.Database, logger log.Logger, iteration models.Iteration) (models.Project, error) {
	project, err := db.GetProject(iteration.ProjectID)
	if err != nil {
		logger.Error(err.Error())
		return project, err
	}

	project.CurrentIteration = iteration
	project.Iterations = append(project.Iterations, iteration)

	err = projects.UpdateProject(db, &project)
	if err != nil {
		logger.Error(err.Error())
		return project, err
	}

	return project, nil
}

// SwapCurrentIteration swaps the current iteration for a project
func SwapCurrentIteration(db database.Database, logger log.Logger, iteration models.Iteration) (models.Project, error) {
	project, err := db.GetProject(iteration.ProjectID)
	if err != nil {
		logger.Error(err.Error())
		return project, err
	}

	valid := checkIfValidIteration(iteration, project.Iterations)
	if !valid {
		e := errors.New(models.InvalidIterationErrorMessage)
		return project, e
	}

	project.CurrentIteration = iteration
	err = projects.UpdateProject(db, &project)
	if err != nil {
		logger.Error(err.Error())
		return project, err
	}

	return project, nil
}

func checkIfValidIteration(iter models.Iteration, list []models.Iteration) bool {
	for _, i := range list {
		if reflect.DeepEqual(i, iter) {
			return true
		}
	}

	return false
}
