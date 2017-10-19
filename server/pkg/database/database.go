package database

import (
	"github.com/darwinfroese/hacksite/server/models"
	"github.com/darwinfroese/hacksite/server/pkg/account"
)

// TODO: Database needs to be a singleton that if it's already
// been created, it should be returned instead of re-created

// Database is an interface to our database needs
type Database interface {
	// Projects
	AddProject(project models.Project) error
	GetProject(id string) (models.Project, error)
	// TODO: UpdateProject - models.Project could probably be removed
	// and the project passed in returned since no internal changes
	// are happening
	UpdateProject(project models.Project) error
	RemoveProject(id string) error

	// Accounts
	CreateAccount(account account.Account) error
	GetAccountByUsername(username string) (account.Account, error)
	GetAccountByEmail(email string) (account.Account, error)
	UpdateAccount(account account.Account) error

	// Sessions
	StoreSession(session models.Session) error
	GetSession(sessionToken string) (models.Session, error)
	GetAllSessions() ([]models.Session, error)
	RemoveSession(sessionToken string) error
}
