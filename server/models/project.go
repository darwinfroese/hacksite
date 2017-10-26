package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

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
