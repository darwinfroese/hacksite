package accounts

import (
	"errors"
	"fmt"
	"reflect"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"

	"github.com/darwinfroese/hacksite/server/models"
	"github.com/darwinfroese/hacksite/server/pkg/auth"
	"github.com/darwinfroese/hacksite/server/pkg/database"
	"github.com/darwinfroese/hacksite/server/pkg/log"
)

// Account contains the information for each user
// type Account struct {
// 	// Username and Email are unique Identifiers
// 	Username, Password, Email, Salt string
// 	ProjectIds                      []string
// }

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

	err = account.ValidateAccount()
	if err != nil {
		return fmt.Errorf("account could not be validated: %s", err.Error())
	}

	account.Password = password
	account.Salt = salt

	err = db.CreateAccount(*account)
	if err != nil {
		return fmt.Errorf("an error occured inserting the account into the database: %s", err.Error())
	}

	return nil
}

//Validate account method
func (account models.Account) ValidateAccount() error {
	return validation.ValidateStruct(&account,
		validation.Field(&account.Username, validation.Required, is.Alphanumeric, validation.Min(3)),
		validation.Field(&account.Email, validation.Required, is.Email),
	)
}
