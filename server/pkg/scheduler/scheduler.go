package scheduler

import (
	"fmt"
	"os"
	"time"

	"github.com/darwinfroese/hacksite/server/pkg/api"
	"github.com/darwinfroese/hacksite/server/pkg/database"
)

// Start executes the scheduler that will run tasks periodically in the background
func Start(ctx *api.Context) {
	ticker := time.NewTicker(5 * time.Minute)

	go func() {
		for range ticker.C {
			// TODO: These should be logs
			count, err := cleanSessions(*ctx.DB)

			if err != nil {
				fmt.Printf("[ERROR] An error occured trying to clean sessions from the database: %s\n", err.Error())
			} else {
				fmt.Printf("[INFO] Removed %d expired sessions\n", count)
			}
		}
	}()
}

// cleanSessions grabs all the sessions from the database and removes the ones that are expired
func cleanSessions(db database.Database) (int, error) {
	count := 0

	sessions, err := db.GetAllSessions()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		return count, err
	}

	for _, sesh := range sessions {
		if time.Now().After(sesh.Expiration) {
			err = db.RemoveSession(sesh.Token)
			if err != nil {
				// since this is just removing one at a time we can continue on if one fails
				fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
			} else {
				count++
			}
		}
	}

	return count, nil
}
