package models

import (
	"time"
)

// TODO: ID values should be scaled up to UINT64 values so that they
// don't roll over at a "low" number

// Project contains a representation of a project
type Project struct {
	ID               int
	Name             string
	Description      string
	Status           string
	CurrentIteration Iteration
	Iterations       []Iteration
}

// Task contains a representation of a task
type Task struct {
	ID              int
	ProjectID       int
	Task            string
	Completed       bool
	IterationNumber int
}

// Iteration contains iteration information for a project
type Iteration struct {
	Number    int
	Tasks     []Task
	ProjectID int
}

// Account contains the information for each user
type Account struct {
	ID                              int
	Username, Password, Email, Salt string
	ProjectIds                      []int
}

// LoginAccount is a simplified account object for login requests
type LoginAccount struct {
	Username, Password string
}

// Session represents the contents of the cookie for the browser
type Session struct {
	Token      string
	UserID     int
	Expiration time.Time
}

// ServerConfig contains server configuration settings
type ServerConfig struct {
	Port, KeyLocation, CertLocation, WebFileLocation string
}

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
