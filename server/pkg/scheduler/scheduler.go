package scheduler

import (
	"fmt"
	"time"

	"github.com/darwinfroese/hacksite/server/pkg/database"
)

// Start executes the scheduler that will run tasks periodically in the background
func Start(db database.Database) {
	ticker := time.NewTicker(5 * time.Minute)

	go func() {
		for range ticker.C {
			// TODO: These should be logs
			count, err := db.CleanSessions()

			if err != nil {
				fmt.Printf("[ERROR] An error occured trying to clean sessions from the database: %s\n", err.Error())
			} else {
				fmt.Printf("[INFO] Removed %d expired sessions\n", count)
			}
		}
	}()
}
