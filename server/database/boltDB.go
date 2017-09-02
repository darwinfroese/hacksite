package database

import (
	"fmt"

	"github.com/darwinfroese/hacksite/server/models"
)

// CreateBoltDB creates a basic database struct
func CreateBoltDB() Database {
	return &boltDB{
		dbLocation: "database.db",
	}
}

// AddProject will add a project to the database
func (db *boltDB) AddProject(project models.Project) {
	fmt.Println("AddProject called")
}

// GetProject will lookup a project by id
func (db *boltDB) GetProject(id int) models.Project {
	fmt.Println("GetProject called. Looking for project", id)
	return models.Project{}
}

// GetProjects will return all the projects in the database
func (db *boltDB) GetProjects() []models.Project {
	fmt.Println("GetProjects called")
	return nil
}

// UpdateTask will update a specfic task in a project
func (db *boltDB) UpdateTask(task models.Task) {
	fmt.Println("UpdateTask called")
}

// RemoveProject will remove a project from the database
func (db *boltDB) RemoveProject(id int) {
	fmt.Println("RemoveProject called")
}

// RemoveTask will remove a task from a project
func (db *boltDB) RemoveTask(task models.Task) {
	fmt.Println("RemoveTask called")
}
