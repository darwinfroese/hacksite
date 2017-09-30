package iterations

import (
	"github.com/darwinfroese/hacksite/server/models"
	"github.com/darwinfroese/hacksite/server/pkg/database"
)

// CreateIteration creates a new iteration and stores it in the database
func CreateIteration(db database.Database, iteration models.Iteration) (models.Project, error) {
	return db.AddIteration(iteration)
}

// SwapCurrentIteration swaps the current iteration for a project
func SwapCurrentIteration(db database.Database, iteration models.Iteration) (models.Project, error) {
	return db.SwapCurrentIteration(iteration)
}
