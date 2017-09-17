package utilities

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"

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

// SaltPassword generates a salt, puts it into the password and returns
// the salt and the new password or an error
func SaltPassword(password string) (string, string, error) {
	salt := make([]byte, sha256.Size)

	n, err := rand.Read(salt)

	if err != nil {
		return "", "", err
	}

	if n != sha256.Size {
		return "", "", err
	}

	saltedVal := append([]byte(password), salt...)
	encrypted := sha256.Sum256(saltedVal)

	hashStr := base64.StdEncoding.EncodeToString(encrypted[:sha256.Size])
	saltStr := base64.StdEncoding.EncodeToString(salt)

	return saltStr, hashStr, nil
}
