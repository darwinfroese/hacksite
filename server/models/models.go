package models

import (
	"time"
)

// TODO: ID values should be scaled up to UINT64 values so that they
// don't roll over at a "low" number

// Project contains a representation of a project
type Project struct {
	ID               string
	Name             string
	Description      string
	Status           string
	CurrentIteration Iteration
	Iterations       []Iteration
}

// Task contains a representation of a task
type Task struct {
	ID              uint64
	ProjectID       string
	Task            string
	Completed       bool
	IterationNumber int
}

// Iteration contains iteration information for a project
type Iteration struct {
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

// LoginAccount is a simplified account object for login requests
type LoginAccount struct {
	Username, Password string
}

// Session represents the contents of the cookie for the browser
type Session struct {
	Token      string
	Username   string
	Expiration time.Time
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
	InvalidIterationErrorMessage = "iteration selected does not exist"
)
