package accounts

import (
	"fmt"
	"os"
	"errors"

	"github.com/darwinfroese/hacksite/server/models"
	"github.com/darwinfroese/hacksite/server/pkg/auth"
	"github.com/darwinfroese/hacksite/server/pkg/database"
)

const (
	// UnathorizedErrorMessage is used to determine if the error is because of unathorized account
	UnathorizedErrorMessage = "Account is invalid"
)

// CreateAccount will create an account and insert it into the database
func CreateAccount(db database.Database, account *models.Account) error {

	//Check if the username already exists
	validUsername, invalidUsername := db.GetAccount(account.Username)
	if invalidUsername != nil {
		//Username is not in use
	} else {
		invalidUsername = errors.New("Username is taken")
		return invalidUsername
	}

	//Check if the email already exists
	validEmail, invalidEmail := db.GetAccountByEmail(account.Email)
	if invalidEmail != nil {
		//Email is not in use
	} else {
		invalidEmail = errors.New("This Email is already in use")
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
