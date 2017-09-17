package models

import (
	"time"
)

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

// Status constants for projects
const (
	StatusCompleted  = "Completed"
	StatusInProgress = "InProgress"
	StatusNew        = "New"
)
