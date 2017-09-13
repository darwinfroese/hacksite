package utilities

import (
	"github.com/darwinfroese/hacksite/server/models"
)

// ValidateProject checks if the model is valid
func ValidateProject(project models.Project) bool {
	return project.Name == ""
}
