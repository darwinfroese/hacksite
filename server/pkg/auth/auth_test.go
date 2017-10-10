package auth

import (
	"os"
	"testing"
	"time"

	"github.com/darwinfroese/hacksite/server/models"
	"github.com/darwinfroese/hacksite/server/pkg/database"
	"github.com/darwinfroese/hacksite/server/pkg/database/bolt"
)

var db database.Database

var username = "test-user"
var password = "secure-password"
var id = 1234

func TestMain(m *testing.M) {
	db = bolt.New()

	// Setup for login_test
	salt, hash, _ := SaltPassword(password)
	db.CreateAccount(models.Account{
		Username: username,
		ID:       id,
		Password: hash,
		Salt:     salt,
	})

	retCode := m.Run()

	os.Remove("database.db")
	os.Exit(retCode)
}

func TestSaltPassword(t *testing.T) {
	t.Log("[ 01 ] Testing SaltPassword does not return the password...")

	salt, hash, err := SaltPassword(password)

	if err != nil {
		t.Errorf("[ FAIL ] There was an error salting the password: %s\n", err.Error())
	}

	if salt == password {
		t.Error("[ FAIL ] The password was returned as the salt value.\n")
	}
	if hash == password {
		t.Error("[ FAIL ] The password was returned as the hash value.\n")
	}
}

func TestGetSaltedPassword(t *testing.T) {
	t.Log("[ 01 ] Testing GetSaltedPassword returns the same password that's generated...")

	password := "secure-password"

	salt, hash, _ := SaltPassword(password)
	hashed, err := GetSaltedPassword(password, salt)

	if err != nil {
		t.Error("[ FAIL ] There was an error getting the salted password.\n")
	}

	if hashed != hash {
		t.Error("[ FAIL ] Did not get the same value out of GetSaltedPassword.\n")
	}
}

func TestCreateSession(t *testing.T) {
	t.Log("[ 01 ] Testing CreateSession should return a session that expires in the future...")

	id := 10

	session := CreateSession(id)

	if time.Now().After(session.Expiration) {
		t.Error("[ FAIL ] The session created already expired.\n")
	}
}
