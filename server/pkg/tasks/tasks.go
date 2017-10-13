package tasks

import (
	"github.com/darwinfroese/hacksite/server/models"
	"github.com/darwinfroese/hacksite/server/pkg/database"
	"github.com/darwinfroese/hacksite/server/pkg/log"
	"github.com/darwinfroese/hacksite/server/pkg/projects"
)

// UpdateTask updates a task in a project and pushes the change into the database
func UpdateTask(db database.Database, logger log.Logger, task models.Task) (models.Project, error) {
	project, err := db.GetProject(task.ProjectID)
	if err != nil {
		logger.Error(err.Error())
		return project, err
	}

	tasks := project.CurrentEvolution.Tasks

	for i, t := range tasks {
		if task.ID == t.ID {
			tasks[i] = task
			break
		}
	}

	project.CurrentEvolution.Tasks = tasks
	for i, iter := range project.Evolutions {
		if iter.Number == project.CurrentEvolution.Number {
			project.Evolutions[i] = project.CurrentEvolution
		}
	}

	err = projects.UpdateProject(db, &project)

	err = db.UpdateProject(project)
	if err != nil {
		logger.Error(err.Error())
		return project, err
	}

	return project, nil
}

// RemoveTask removes a task from a project and pushes the change into the database
func RemoveTask(db database.Database, logger log.Logger, task models.Task) (models.Project, error) {
	project, err := db.GetProject(task.ProjectID)
	if err != nil {
		logger.Error(err.Error())
		return project, err
	}

	tasks := project.CurrentEvolution.Tasks

	for i, t := range tasks {
		if task.ID == t.ID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}

	// updates both current evolution as well as the same evolution in the
	// evolutions list
	project.CurrentEvolution.Tasks = tasks
	for i, iter := range project.Evolutions {
		if iter.Number == project.CurrentEvolution.Number {
			project.Evolutions[i] = project.CurrentEvolution
		}
	}

	err = projects.UpdateProject(db, &project)
	if err != nil {
		logger.Error(err.Error())
		return project, err
	}

	return project, nil
}
