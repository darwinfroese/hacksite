package models

import (
	"time"
)

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

// Session represents the contents of the cookie for the browser
type Session struct {
	Token      string
	Username   string
	Expiration time.Time
	RememberMe bool
}
