package scheduler

import (
	"os"
	"testing"
	"time"

	"github.com/darwinfroese/hacksite/server/models"
	"github.com/darwinfroese/hacksite/server/pkg/database"
	"github.com/darwinfroese/hacksite/server/pkg/database/bolt"
	"github.com/darwinfroese/hacksite/server/pkg/log"
	"github.com/darwinfroese/hacksite/server/pkg/log/testLogger"
)

var expiredSession = models.Session{
	Token:      "TestSession1",
	Username:   "test-account",
	Expiration: time.Now().Add(time.Duration(-1) * time.Hour),
}
var unexpiredSession = models.Session{
	Token:      "TestSession2",
	Username:   "test-account",
	Expiration: time.Now().Add(time.Duration(10) * time.Hour),
}

var db database.Database
var logger log.Logger

func TestMain(m *testing.M) {
	db = bolt.New()
	logger = testLogger.New()
	db.StoreSession(expiredSession)
	db.StoreSession(unexpiredSession)

	retCode := m.Run()

	os.Remove("database.db")
	os.Exit(retCode)
}

var sessionTests = []struct {
	Description          string
	ExpectedRemovedCount int
}{{
	Description:          "Calling clean sessions should only remove the expired sessions.",
	ExpectedRemovedCount: 1,
}, {
	Description:          "Calling clean sessions should remove no sessions if there are no expired sessions.",
	ExpectedRemovedCount: 0,
}}

func TestCleanSessions(t *testing.T) {
	t.Log("Testing CleanSessions...")

	for i, tc := range sessionTests {
		t.Logf("[ %02d ] %s\n", i+1, tc.Description)

		count, err := cleanSessions(db, logger)
		if err != nil {
			t.Errorf("[ FAIL ] An error occured trying to remove expried sessions: %s\n", err.Error())
		}

		if count != tc.ExpectedRemovedCount {
			t.Errorf("[ FAIL ] Expected to remove %d sessions but removed %d sessions.", tc.ExpectedRemovedCount, count)
		}
	}
}
