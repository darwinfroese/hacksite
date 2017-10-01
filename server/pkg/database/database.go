package database

import "github.com/darwinfroese/hacksite/server/models"

// TODO: Database needs to be a singleton that if it's already
// been created, it should be returned instead of re-created

// Database is an interface to our database needs
type Database interface {
	// Projects
	AddProject(project models.Project) (models.Project, error)
	GetProject(id int) (models.Project, error)
	GetProjects(userID int) ([]models.Project, error)
	GetNextProjectID() (int, error)
	// TODO: UpdateProject - models.Project could probably be removed
	// and the project passed in returned since no internal changes
	// are happening
	UpdateProject(project models.Project) error
	RemoveProject(id int) error

	// Tasks
	UpdateTask(task models.Task) (models.Project, error)
	RemoveTask(task models.Task) (models.Project, error)

	// Iterations
	AddIteration(iteration models.Iteration) (models.Project, error)
	SwapCurrentIteration(iteration models.Iteration) (models.Project, error)

	// Accounts
	CreateAccount(account models.Account) (int, error)
	GetAccount(username string) (models.Account, error)
	GetAccountByID(userID int) (models.Account, error)
	GetNextAccountID() (int, error)
	UpdateAccount(account models.Account) error

	// Sessions
	StoreSession(session models.Session) error
	GetSession(sessionToken string) (models.Session, error)
	GetAllSessions() ([]models.Session, error)
	GetNextSessionID() (int, error)
	CleanSessions() (int, error)
	RemoveSession(sessionToken string) error
}

type boltDB struct {
	dbLocation string
}

// CreateDB returns an instance of the database
// depending on the environment
func CreateDB() Database {
	return createBoltDB()
}
