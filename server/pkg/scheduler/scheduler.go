package scheduler

import (
	"fmt"
	"time"

	"github.com/darwinfroese/hacksite/server/pkg/api"
	"github.com/darwinfroese/hacksite/server/pkg/database"
	"github.com/darwinfroese/hacksite/server/pkg/log"
)

// Start executes the scheduler that will run tasks periodically in the background
func Start(ctx api.Context) {
	ticker := time.NewTicker(5 * time.Minute)

	go func() {
		for range ticker.C {
			// TODO: These should be logs
			count, err := cleanSessions(*ctx.DB, *ctx.Logger)

			if err != nil {
				(*ctx.Logger).Error(fmt.Sprintf("Attempting to clean sessions from database: %s", err.Error()))
			} else {
				(*ctx.Logger).Info(fmt.Sprintf("Removed %d expired sessions", count))
			}
		}
	}()
}

// cleanSessions grabs all the sessions from the database and removes the ones that are expired
func cleanSessions(db database.Database, logger log.Logger) (int, error) {
	count := 0

	sessions, err := db.GetAllSessions()
	if err != nil {
		logger.Error(err.Error())
		return count, err
	}

	for _, sesh := range sessions {
		if time.Now().After(sesh.Expiration) && sesh.RememberMe != true {
			err = db.RemoveSession(sesh.Token)
			if err != nil {
				// since this is just removing one at a time we can continue on if one fails
				logger.Error(err.Error())
			} else {
				count++
			}
		}
	}

	return count, nil
}
