package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// Account contains the information for each user
type Account struct {
	// Username and Email are unique Identifiers
	Username, Password, Name, Email, Salt string
	ProjectIds                            []string
}

//Validate account method
func (account Account) Validate() error {
	return validation.ValidateStruct(&account,
		validation.Field(&account.Username, validation.Required, is.Alphanumeric, validation.Length(3, 50)),
		validation.Field(&account.Email, validation.Required, is.Email),
	)
}
