package accounts

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/darwinfroese/hacksite/server/models"
	"github.com/darwinfroese/hacksite/server/pkg/auth"
	"github.com/darwinfroese/hacksite/server/pkg/database"
	"github.com/darwinfroese/hacksite/server/pkg/log"
)

const (
	invalidAccountFormatter   = "create account request is missing: %s"
	usernameTakenErrorMessage = "username is already taken"
	emailTakenErrorMessage    = "this email is already in use"
)

// CreateAccount will create an account and insert it into the database
func CreateAccount(db database.Database, logger log.Logger, account *models.Account) *models.APIError {

	//Check if the username already exists
	acc, err := db.GetAccountByUsername(account.Username)
	if err != nil {
		logger.Error(fmt.Sprintf("Error getting account: %s", err.Error()))
		return &models.APIError{Message: err.Error(), Code: http.StatusInternalServerError}
	}
	if !reflect.DeepEqual(acc, (models.Account{})) {
		return &models.APIError{Message: usernameTakenErrorMessage, Code: http.StatusConflict}
	}

	//Check if the email already exists
	acc, err = db.GetAccountByEmail(account.Email)
	if err != nil {
		logger.Error(fmt.Sprintf("Error getting account: %s\n", err.Error()))
		return &models.APIError{Message: err.Error(), Code: http.StatusInternalServerError}
	}
	if !reflect.DeepEqual(acc, (models.Account{})) {
		return &models.APIError{Message: emailTakenErrorMessage, Code: http.StatusConflict}
	}

	salt, password, err := auth.SaltPassword(account.Password)
	if err != nil {
		return &models.APIError{
			Message: fmt.Sprintf("an error occured salting the account password: %s", err.Error()),
			Code:    http.StatusInternalServerError,
		}
	}

	err = account.Validate()
	if err != nil {
		return &models.APIError{
			Message: fmt.Sprintf("account could not be validated: %s", err.Error()),
			Code:    http.StatusBadRequest,
		}
	}

	account.Password = password
	account.Salt = salt

	err = db.CreateAccount(*account)
	if err != nil {
		return &models.APIError{
			Message: fmt.Sprintf("an error occured inserting the account into the database: %s", err.Error()),
			Code:    http.StatusInternalServerError,
		}
	}

	return nil
}

// UpdateAccount removes the existing account and inserts a new one into the database
func UpdateAccount(db database.Database, username, email string, newAccount models.Account) *models.APIError {
	err := newAccount.Validate()
	if err != nil {
		return &models.APIError{
			Message: fmt.Sprintf("account could not be validated: %s", err.Error()),
			Code:    http.StatusBadRequest,
		}
	}

	if username != newAccount.Username {
		//Check if the username already exists
		acc, err := db.GetAccountByUsername(newAccount.Username)
		if err != nil {
			return &models.APIError{Message: err.Error(), Code: http.StatusInternalServerError}
		}
		if !reflect.DeepEqual(acc, (models.Account{})) {
			return &models.APIError{Message: usernameTakenErrorMessage, Code: http.StatusConflict}
		}
	}

	if email != newAccount.Email {
		//Check if the email already exists
		acc, err := db.GetAccountByEmail(newAccount.Email)
		if err != nil {
			return &models.APIError{Message: err.Error(), Code: http.StatusInternalServerError}
		}
		if !reflect.DeepEqual(acc, (models.Account{})) {
			return &models.APIError{Message: emailTakenErrorMessage, Code: http.StatusConflict}
		}
	}

	err = db.RemoveAccount(username, email)
	if err != nil {
		return &models.APIError{
			DeveloperMessage: fmt.Sprintf("There was an error removing the existing account to be updated: %s", err.Error()),
			Message:          fmt.Sprint("account could not be updated"),
			Code:             http.StatusInternalServerError,
		}
	}

	err = db.CreateAccount(newAccount)
	if err != nil {
		return &models.APIError{
			DeveloperMessage: fmt.Sprintf("There was an error inserting the new account: %s", err.Error()),
			Message:          fmt.Sprint("account could not be updated"),
			Code:             http.StatusInternalServerError,
		}
	}

	return nil
}
