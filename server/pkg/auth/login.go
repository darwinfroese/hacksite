package auth

import (
	"fmt"

	"github.com/darwinfroese/hacksite/server/models"
	"github.com/darwinfroese/hacksite/server/pkg/database"
)

const (
	// UnathorizedErrorMessage is used to determine if the error is because of unathorized account
	UnathorizedErrorMessage = "Account is invalid"
)

// Login will attempt to login the user and return a new session
func Login(db database.Database, username, password string) (models.Session, error) {
	account, err := db.GetAccountByUsername(username)
	if err != nil {
		if err.Error() == "no matching account found" {
			return models.Session{}, fmt.Errorf("Error: %s", UnathorizedErrorMessage)
		}

		return models.Session{}, fmt.Errorf("There was a problem getting account information: %s", err.Error())
	}

	password, err = GetSaltedPassword(password, account.Salt)
	if err != nil {
		return models.Session{}, fmt.Errorf("There was a problem salting the password: %s", err.Error())
	}

	if password != account.Password {
		return models.Session{}, fmt.Errorf("Error: %s", UnathorizedErrorMessage)
	}

	session := CreateSession(account.ID)
	err = db.StoreSession(session)

	if err != nil {
		return models.Session{}, fmt.Errorf("There was a problem storing the session: %s", err.Error())
	}

	return session, nil
}
