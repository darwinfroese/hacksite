package database

import "github.com/darwinfroese/hacksite/server/models"

type Database interface {
	AddProject(project models.Project)
	GetProject(id int) models.Project
	GetProjects() []models.Project
	UpdateTask(task models.Task)
	RemoveProject(id int)
	RemoveTask(task models.Task)
}

type boltDB struct {
	dbLocation string
}
