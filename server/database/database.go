package database

import "github.com/darwinfroese/hacksite/server/models"

// Database is an interface to our database needs
type Database interface {
	AddProject(project models.Project) (models.Project, error)
	GetProject(id int) (models.Project, error)
	GetProjects() ([]models.Project, error)
	UpdateTask(task models.Task) (models.Project, error)
	RemoveProject(id int) error
	RemoveTask(task models.Task) (models.Project, error)
}

type boltDB struct {
	dbLocation string
}
