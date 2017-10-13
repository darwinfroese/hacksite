package evolutions

import (
	"errors"
	"reflect"

	"github.com/darwinfroese/hacksite/server/models"
	"github.com/darwinfroese/hacksite/server/pkg/database"
	"github.com/darwinfroese/hacksite/server/pkg/log"
	"github.com/darwinfroese/hacksite/server/pkg/projects"
)

// CreateEvolution creates a new evolution and stores it in the database
func CreateEvolution(db database.Database, logger log.Logger, evolution models.Evolution) (models.Project, error) {
	project, err := db.GetProject(evolution.ProjectID)
	if err != nil {
		logger.Error(err.Error())
		return project, err
	}

	project.CurrentEvolution = evolution
	project.Evolutions = append(project.Evolutions, evolution)

	err = projects.UpdateProject(db, &project)
	if err != nil {
		logger.Error(err.Error())
		return project, err
	}

	return project, nil
}

// SwapCurrentEvolution swaps the current evolution for a project
func SwapCurrentEvolution(db database.Database, logger log.Logger, evolution models.Evolution) (models.Project, error) {
	project, err := db.GetProject(evolution.ProjectID)
	if err != nil {
		logger.Error(err.Error())
		return project, err
	}

	valid := checkIfValidEvolution(evolution, project.Evolutions)
	if !valid {
		e := errors.New(models.InvalidEvolutionErrorMessage)
		return project, e
	}

	project.CurrentEvolution = evolution
	err = projects.UpdateProject(db, &project)
	if err != nil {
		logger.Error(err.Error())
		return project, err
	}

	return project, nil
}

func checkIfValidEvolution(iter models.Evolution, list []models.Evolution) bool {
	for _, i := range list {
		if reflect.DeepEqual(i, iter) {
			return true
		}
	}

	return false
}
