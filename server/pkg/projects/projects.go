package projects

import "github.com/darwinfroese/hacksite/server/models"

// UpdateProjectStatus returns the string representation of
// the project's status
func UpdateProjectStatus(project models.Project) string {
	complete := 0
	status := models.StatusNew

	tasks := project.CurrentIteration.Tasks
	for _, task := range tasks {
		if task.Completed {
			complete++
		}
	}
	if complete == len(tasks) {
		status = models.StatusCompleted
	} else if complete > 0 {
		status = models.StatusInProgress
	} else {
		status = models.StatusNew
	}

	return status
}
