package tasks

import (
	"os"
	"testing"

	"github.com/darwinfroese/hacksite/server/pkg/database"
)

var db database.Database

func TestMain(m *testing.M) {
	db = database.CreateDB()

	retCode := m.RunTests()

	os.Remove("database.db")
	os.Exit(retCode)
}

func TestUpdateTask(t *testing.T) {
	t.Fail()
}
