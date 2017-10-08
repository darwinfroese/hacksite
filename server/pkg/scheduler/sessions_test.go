package scheduler

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/darwinfroese/hacksite/server/pkg/database"

	"github.com/darwinfroese/hacksite/server/models"
)

var expiredSession = models.Session{
	Token:      "TestSession1",
	UserID:     1234,
	Expiration: time.Now().Add(time.Duration(-1) * time.Hour),
}
var unexpiredSession = models.Session{
	Token:      "TestSession2",
	UserID:     1234,
	Expiration: time.Now().Add(time.Duration(10) * time.Hour),
}

var db database.Database

func TestMain(m *testing.M) {
	db = database.CreateDB()
	db.StoreSession(expiredSession)
	db.StoreSession(unexpiredSession)

	sessions, err := db.GetAllSessions()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(sessions)

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

		count, err := cleanSessions(db)
		if err != nil {
			t.Errorf("[ FAIL ] An error occured trying to remove expried sessions: %s\n", err.Error())
		}

		if count != tc.ExpectedRemovedCount {
			t.Errorf("[ FAIL ] Expected to remove %d sessions but removed %d sessions.", tc.ExpectedRemovedCount, count)
		}
	}
}
