package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// TODO: ID values should be scaled up to UINT64 values so that they
// don't roll over at a "low" number

// Project contains a representation of a project
type Project struct {
	ID               string
	Name             string
	Description      string
	Status           string
	CurrentEvolution Evolution
	Evolutions       []Evolution
}

// Validate checks if the model is valid
func (project Project) Validate() error {
	return validation.ValidateStruct(&project,
		validation.Field(&project.Name, validation.Required),
	)
}

// Task contains a representation of a task
type Task struct {
	ID              uint64
	ProjectID       string
	Task            string
	Completed       bool
	EvolutionNumber int
}

// Evolution contains evolution information for a project
type Evolution struct {
	Number    int
	Tasks     []Task
	ProjectID string
}

// Account contains the information for each user
type Account struct {
	// Username and Email are unique Identifiers
	Username, Password, Email, Salt string
	ProjectIds                      []string
}

//Validate account method
func (account Account) Validate() error {
	return validation.ValidateStruct(&account,
		validation.Field(&account.Username, validation.Required, is.Alphanumeric, validation.Length(3, 50)),
		validation.Field(&account.Email, validation.Required, is.Email),
	)
}

// LoginAccount is a simplified account object for login requests
type LoginAccount struct {
	Username, Password string
}

// Session represents the contents of the cookie for the browser
type Session struct {
	Token      string
	Username   string
	Expiration time.Time
	RememberMe bool
}

// ResponseObject is a wrapper for responding to error requests
type ResponseObject struct {
	StatusCode   int
	ErrorMessage string
	Message      string
}

// Status constants for projects
const (
	StatusCompleted              = "Completed"
	StatusInProgress             = "InProgress"
	StatusNew                    = "New"
	UsernameTakenErrorMessage    = "username is already taken"
	EmailTakenErrorMessage       = "this email is already in use"
	InvalidEvolutionErrorMessage = "evolution selected does not exist"
)
