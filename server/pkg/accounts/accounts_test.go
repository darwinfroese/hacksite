package accounts

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/darwinfroese/hacksite/server/models"
	"github.com/darwinfroese/hacksite/server/pkg/database"
	"github.com/darwinfroese/hacksite/server/pkg/database/bolt"
	"github.com/darwinfroese/hacksite/server/pkg/log"
	"github.com/darwinfroese/hacksite/server/pkg/log/testLogger"
)

var db database.Database
var logger log.Logger

// TestMain lets us setup the database and then remove it when we are done
func TestMain(m *testing.M) {
	db = bolt.New()
	logger = testLogger.New()

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
}, {
	Description: "Attempting to create an account with an username already in use shoud fail.",
	Account: models.Account{
		Username: "test-account",
		Password: "secure-password",
		Email:    "testemail@email.com",
	},
	ExpectedAccount: models.Account{
		Username: "test-account",
		Password: "secure-password",
		Email:    "testemail@email.com",
	},
	ExpectedErrorMessage: models.UsernameTakenErrorMessage,
}, {
	Description: "Attempting to create an account with an email already in use should fail.",
	Account: models.Account{
		Username: "account-test",
		Password: "secure-password",
		Email:    "test@email.com",
	},
	ExpectedAccount: models.Account{
		Username: "account-test",
		Password: "secure-password",
		Email:    "test@email.com",
	},
	ExpectedErrorMessage: models.EmailTakenErrorMessage,
}}

func TestCreateAccount(t *testing.T) {
	t.Logf("Starting CreateAccount Tests...")

	for i, tc := range createAccountTests {
		t.Logf("[ %02d ] %s\n", i+1, tc.Description)
		err := CreateAccount(db, logger, &tc.Account)

		if err != nil && err.Error() != tc.ExpectedErrorMessage {
			t.Errorf("[ FAIL ] The test failed because of an unexpected error: %s\n", err.Error())
			// go to the next test case since we want to still loop through all of them
			break
		}

		// These have to be copied over or deep equal will fail
		tc.ExpectedAccount.Password = tc.Account.Password
		tc.ExpectedAccount.Salt = tc.Account.Salt

		if !reflect.DeepEqual(tc.Account, tc.ExpectedAccount) {
			t.Errorf("[ FAIL ] CreateAccount did not return the account expected.\nExpected: %v\nbut got:  %v\n",
				tc.Account, tc.ExpectedAccount)
		}
	}
}

//Validator TEST
var validateAccountTests = []struct {
	Username    string
	Email       string
	ExpectedResult bool
}{{
	Username: "alimasyhur",
	Email:	"jegrag4ever@gmail.com"
	ExpectedResult: true,
}, {
	Username:    "am",
	Email	: "jegrag4ever"
	ExpectedResult: false,
}}

func TestValidateAccount(t *testing.T) {
	t.Log("Testing ValidateAccount...")

	for i, tc := range validateAccountTests {
		t.Logf("[ %02d ] %s\n", i+1, tc.Description)

		result := tc.Validate()
		if result != tc.ExpectedResult {
			t.Errorf("[ FAIL ] ValidateProject did not return expected value. Expected \"%v\" but got \"%v\".",
				result, tc.ExpectedResult)
		}
	}
}
