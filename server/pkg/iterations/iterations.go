package iterations

import (
	"fmt"
	"os"

	"github.com/darwinfroese/hacksite/server/pkg/projects"

	"github.com/darwinfroese/hacksite/server/models"
	"github.com/darwinfroese/hacksite/server/pkg/database"
)

// CreateIteration creates a new iteration and stores it in the database
func CreateIteration(db database.Database, iteration models.Iteration) (models.Project, error) {
	project, err := db.GetProject(iteration.ProjectID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		return project, err
	}

	project.CurrentIteration = iteration
	project.Iterations = append(project.Iterations, iteration)

	err = projects.UpdateProject(db, &project)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		return project, err
	}

	return project, nil
}

// SwapCurrentIteration swaps the current iteration for a project
func SwapCurrentIteration(db database.Database, iteration models.Iteration) (models.Project, error) {
	project, err := db.GetProject(iteration.ProjectID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		return project, err
	}

	project.CurrentIteration = iteration
	err = projects.UpdateProject(db, &project)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		return project, err
	}

	return project, nil
}
