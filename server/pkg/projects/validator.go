package projects

import (
	"github.com/darwinfroese/hacksite/server/models"
	validation "src/github.com/go-ozzo/ozzo-validation"
)

// ValidateProject checks if the model is valid
func (project models.Project) ValidateProject() error {
	return validation.ValidateStruct(&project,
		validation.Field(&project.Name, validation.Required),
	)
}
