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
	err := validation.ValidateStruct(&project,
		validation.Field(&project.Name, validation.Required, validation.Length(1, 32)),
		validation.Field(&project.Description, validation.Length(0, 128)),
	)

	if err != nil {
		return err
	}

	err = validateEvolution(project.CurrentEvolution)
	if err != nil {
		return err
	}

	return nil
}

func validateEvolution(evo Evolution) error {
	for _, task := range evo.Tasks {
		err := validation.ValidateStruct(&task,
			validation.Field(&task.Task, validation.Length(1, 32)),
		)

		if err != nil {
			return err
		}
	}

	return nil
}
