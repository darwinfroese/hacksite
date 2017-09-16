package models

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

// User contains the information for each user
type User struct {
	Name, Password, Email string
	confirmPassword, salt string
	ProjectIds            []int
}

// Status constants for projects
const (
	StatusCompleted  = "Completed"
	StatusInProgress = "InProgress"
	StatusNew        = "New"
)
