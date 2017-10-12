package accounts

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/darwinfroese/hacksite/server/models"
	"github.com/darwinfroese/hacksite/server/pkg/auth"
	"github.com/darwinfroese/hacksite/server/pkg/database"
	"github.com/darwinfroese/hacksite/server/pkg/log"
)

const (
	invalidAccountFormatter = "create account request is missing: %s"
)

// CreateAccount will create an account and insert it into the database
func CreateAccount(db database.Database, logger log.Logger, account *models.Account) error {

	//Check if the username already exists
	acc, err := db.GetAccountByUsername(account.Username)
	if err != nil {
		logger.Error(fmt.Sprintf("Error getting account: %s", err.Error()))
		return err
	}
	if !reflect.DeepEqual(acc, (models.Account{})) {
		return errors.New(models.UsernameTakenErrorMessage)
	}

	//Check if the email already exists
	acc, err = db.GetAccountByEmail(account.Email)
	if err != nil {
		logger.Error(fmt.Sprintf("Error getting account: %s\n", err.Error()))
		return err
	}
	if !reflect.DeepEqual(acc, (models.Account{})) {
		return errors.New(models.EmailTakenErrorMessage)
	}

	salt, password, err := auth.SaltPassword(account.Password)
	if err != nil {
		return fmt.Errorf("an error occured salting the account password: %s", err.Error())
	}

	err = validateAccount(*account)
	if err != nil {
		return fmt.Errorf("account could not be validated: %s", err.Error())
	}

	account.Password = password
	account.Salt = salt
	id, err := db.GetNextAccountID()
	if err != nil {
		logger.Error(fmt.Sprintf("Error: %s\n", err.Error()))
		return err
	}

	account.ID = id
	_, err = db.CreateAccount(*account)
	if err != nil {
		return fmt.Errorf("an error occured inserting the account into the database: %s", err.Error())
	}

	return nil
}

func validateAccount(account models.Account) error {
	if account.Email == "" {
		return fmt.Errorf(invalidAccountFormatter, "email")
	}
	if account.Username == "" {
		return fmt.Errorf(invalidAccountFormatter, "username")
	}

	return nil
}
