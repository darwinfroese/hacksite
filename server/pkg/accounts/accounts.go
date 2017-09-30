package accounts

import (
	"fmt"

	"github.com/darwinfroese/hacksite/server/models"
	"github.com/darwinfroese/hacksite/server/pkg/auth"
	"github.com/darwinfroese/hacksite/server/pkg/database"
)

// CreateAccount will create an account and insert it into the database
func CreateAccount(db database.Database, account models.Account) error {
	salt, password, err := auth.SaltPassword(account.Password)
	if err != nil {
		return fmt.Errorf("An error occured salting the account password: %s", err.Error())
	}

	account.Password = password
	account.Salt = salt

	_, err = db.CreateAccount(account)

	if err != nil {
		return fmt.Errorf("An error occured inserting the account into the database: %s", err.Error())
	}

	return nil
}
