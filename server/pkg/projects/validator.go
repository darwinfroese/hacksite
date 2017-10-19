package projects

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// ValidateProject checks if the model is valid
func (project Project) ValidateProject() error {
	return validation.ValidateStruct(&project,
		validation.Field(&project.Name, validation.Required),
	)
}
