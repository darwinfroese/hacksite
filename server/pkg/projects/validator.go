package projects

import (
	"github.com/darwinfroese/hacksite/server/models"
)

// ValidateProject checks if the model is valid
func ValidateProject(project models.Project) bool {
	// TODO: Find a cleaner validation system
	valid := project.Name == ""

	return !valid
}
