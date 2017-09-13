package utilities

import (
	"github.com/darwinfroese/hacksite/server/models"
)

// ValidateProject checks if the model is valid
func ValidateProject(project models.Project) bool {
	// TODO: Find a cleaner validation system
	valid := project.Name == ""

	return !valid
}

// UpdateProjectStatus returns the string representation of
// the project's status
func UpdateProjectStatus(project models.Project) string {
	complete := 0
	status := models.StatusNew

	for _, task := range project.Tasks {
		if task.Completed {
			complete++
		}
	}
	if complete == len(project.Tasks) {
		status = models.StatusCompleted
	} else if complete > 0 {
		status = models.StatusInProgress
	} else {
		status = models.StatusNew
	}

	return status
}
