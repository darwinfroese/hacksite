package database

import "github.com/darwinfroese/hacksite/server/models"

// Database is an interface to our database needs
type Database interface {
	// Projects
	AddProject(project models.Project) (models.Project, error)
	GetProject(id int) (models.Project, error)
	GetProjects() ([]models.Project, error)
	// TODO: UpdateProject - models.Project could probably be removed
	// and the project passed in returned since no internal changes
	// are happening
	UpdateProject(models.Project) error
	RemoveProject(id int) error

	// Tasks
	UpdateTask(task models.Task) (models.Project, error)
	RemoveTask(task models.Task) (models.Project, error)

	// Iterations
	AddIteration(models.Iteration) (models.Project, error)
	SwapCurrentIteration(models.Iteration) (models.Project, error)

	// Users
	CreateUser(models.User) error
}

type boltDB struct {
	dbLocation string
}

// CreateDB returns an instance of the database
// depending on the environment
func CreateDB() Database {
	return createBoltDB()
}
