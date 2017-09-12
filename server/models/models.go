package models

// Project contains a representation of
// a project
type Project struct {
	ID          int
	Name        string
	Description string
	Tasks       []Task
	Completed   bool
}

// Task contains a representation of a task
type Task struct {
	ID        int
	ProjectID int
	Task      string
	Completed bool
}
