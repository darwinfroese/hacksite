package accounts

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/darwinfroese/hacksite/server/pkg/database"

	"github.com/darwinfroese/hacksite/server/models"
)

var db database.Database

// TestMain lets us setup the database and then remove it when we are done
func TestMain(m *testing.M) {
	db = database.CreateDB()

	retCode := m.Run()

	os.Remove("./database.db")
	os.Exit(retCode)
}

var createAccountTests = []struct {
	Description              string
	Account, ExpectedAccount models.Account
	ExpectedErrorMessage     string
}{{
	Description: "Creating an account should return a valid account model.",
	Account: models.Account{
		Username: "test-account",
		Password: "secure-password",
		Email:    "test@email.com",
	},
	ExpectedAccount: models.Account{
		Username: "test-account",
		Password: "secure-password",
		Email:    "test@email.com",
		ID:       1,
	},
	ExpectedErrorMessage: "",
}, {
	Description: "Creating a second account should increment the ID by one.",
	Account: models.Account{
		Username: "test-account2",
		Password: "secure-password",
		Email:    "test2@email.com",
	},
	ExpectedAccount: models.Account{
		Username: "test-account2",
		Password: "secure-password",
		Email:    "test2@email.com",
		ID:       2,
	},
	ExpectedErrorMessage: "",
}, {
	Description: "Attempting to create an account without an email should fail.",
	Account: models.Account{
		Username: "test-account3",
		Password: "secure-password",
	},
	ExpectedAccount: models.Account{
		Username: "test-account3",
		Password: "secure-password",
	},
	ExpectedErrorMessage: "account could not be validated: " + fmt.Sprintf(invalidAccountFormatter, "email"),
}, {
	Description: "Attempting to create an account without an username should fail.",
	Account: models.Account{
		Password: "secure-password",
		Email:    "test3@email.com",
	},
	ExpectedAccount: models.Account{
		Password: "secure-password",
		Email:    "test3@email.com",
	},
	ExpectedErrorMessage: "account could not be validated: " + fmt.Sprintf(invalidAccountFormatter, "username"),
}}

func TestCreateAccount(t *testing.T) {
	fmt.Println("Starting CreateAccount Tests...")

	for i, tc := range createAccountTests {
		fmt.Printf("[ %02d ] %s\n", i+1, tc.Description)
		err := CreateAccount(db, &tc.Account)

		if err != nil && err.Error() != tc.ExpectedErrorMessage {
			fmt.Printf("[ FAIL ] The test failed because of an unexpected error: %s\n", err.Error())
			t.Fail()
			// go to the next test case since we want to still loop through all of them
			break
		}

		// These have to be copied over or deep equal will fail
		tc.ExpectedAccount.Password = tc.Account.Password
		tc.ExpectedAccount.Salt = tc.Account.Salt

		if !reflect.DeepEqual(tc.Account, tc.ExpectedAccount) {
			fmt.Printf("[ FAIL ] CreateAccount did not return the account expected.\nExpected: %v\nbut got:  %v\n",
				tc.Account, tc.ExpectedAccount)
			t.Fail()
		}
	}
}
