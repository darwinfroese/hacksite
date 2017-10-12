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

	tasks := project.CurrentIteration.Tasks

	for i, t := range tasks {
		if task.ID == t.ID {
			tasks[i] = task
			break
		}
	}

	project.CurrentIteration.Tasks = tasks
	for i, iter := range project.Iterations {
		if iter.Number == project.CurrentIteration.Number {
			project.Iterations[i] = project.CurrentIteration
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

	tasks := project.CurrentIteration.Tasks

	for i, t := range tasks {
		if task.ID == t.ID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}

	// updates both current iteration as well as the same iteration in the
	// iterations list
	project.CurrentIteration.Tasks = tasks
	for i, iter := range project.Iterations {
		if iter.Number == project.CurrentIteration.Number {
			project.Iterations[i] = project.CurrentIteration
		}
	}

	err = projects.UpdateProject(db, &project)
	if err != nil {
		logger.Error(err.Error())
		return project, err
	}

	return project, nil
}
