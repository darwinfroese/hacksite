package accounts

import (
	"fmt"
	"os"

	"github.com/darwinfroese/hacksite/server/models"
	"github.com/darwinfroese/hacksite/server/pkg/auth"
	"github.com/darwinfroese/hacksite/server/pkg/database"
)

// CreateAccount will create an account and insert it into the database
func CreateAccount(db database.Database, account *models.Account) error {

	//Check if the username already exists
	_, invalidUsername := db.GetAccountByUsername(account.Username)
	if invalidUsername == nil {
		invalidUsername = fmt.Errorf(models.UsernameTakenErrorMessage)
		return invalidUsername
	}

	//Check if the email already exists
	_, invalidEmail := db.GetAccountByEmail(account.Email)
	if invalidEmail == nil {
		invalidEmail = fmt.Errorf(models.EmailTakenErrorMessage)
		return invalidEmail
	}

	salt, password, err := auth.SaltPassword(account.Password)
	if err != nil {
		return fmt.Errorf("An error occured salting the account password: %s", err.Error())
	}

	account.Password = password
	account.Salt = salt
	id, err := db.GetNextAccountID()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		return err
	}

	account.ID = id
	_, err = db.CreateAccount(*account)
	if err != nil {
		return fmt.Errorf("An error occured inserting the account into the database: %s", err.Error())
	}

	return nil
}
