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
		Username: "testAccount",
		Password: "secure-password",
		Name:     "Test",
		Email:    "test@email.com",
	},
	ExpectedAccount: models.Account{
		Username: "testAccount",
		Password: "secure-password",
		Name:     "Test",
		Email:    "test@email.com",
	},
	ExpectedErrorMessage: "",
}, {
	Description: "Attempting to create an account without an email should fail.",
	Account: models.Account{
		Username: "testAccount3",
		Password: "secure-password",
		Name:     "test account",
	},
	ExpectedAccount: models.Account{
		Username: "testAccount3",
		Password: "secure-password",
		Name:     "test account",
	},
	ExpectedErrorMessage: "account could not be validated: Email: cannot be blank.",
}, {
	Description: "Attempting to create an account without an username should fail.",
	Account: models.Account{
		Password: "secure-password",
		Email:    "test3@email.com",
		Name:     "test account",
	},
	ExpectedAccount: models.Account{
		Password: "secure-password",
		Email:    "test3@email.com",
		Name:     "test account",
	},
	ExpectedErrorMessage: "account could not be validated: Username: cannot be blank.",
}, {
	Description: "Attempting to create an account without a name should fail.",
	Account: models.Account{
		Password: "secure-password",
		Email:    "test3@email.com",
		Username: "testaccount",
	},
	ExpectedAccount: models.Account{
		Password: "secure-password",
		Email:    "test3@email.com",
		Username: "testaccount",
	},
	ExpectedErrorMessage: "account could not be validated: Name: cannot be blank.",
}, {
	Description: "Attempting to create an account with an username already in use shoud fail.",
	Account: models.Account{
		Username: "testAccount",
		Password: "secure-password",
		Email:    "testemail@email.com",
	},
	ExpectedAccount: models.Account{
		Username: "testAccount",
		Password: "secure-password",
		Email:    "testemail@email.com",
	},
	ExpectedErrorMessage: usernameTakenErrorMessage,
}, {
	Description: "Attempting to create an account with an email already in use should fail.",
	Account: models.Account{
		Username: "testAccount123",
		Password: "secure-password",
		Email:    "test@email.com",
		Name:     "account name",
	},
	ExpectedAccount: models.Account{
		Username: "testAccount123",
		Password: "secure-password",
		Email:    "test@email.com",
		Name:     "account name",
	},
	ExpectedErrorMessage: emailTakenErrorMessage,
}}

func TestCreateAccount(t *testing.T) {
	t.Logf("Starting CreateAccount Tests...")

	for i, tc := range createAccountTests {
		t.Logf("[ %02d ] %s\n", i+1, tc.Description)
		err := CreateAccount(db, logger, &tc.Account)

		if err != nil && err.Error() != tc.ExpectedErrorMessage {
			fmt.Println("err: ", err.Error())
			fmt.Println("tc.ExpectedErrorMessage", tc.ExpectedErrorMessage)
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
	Description    string
	Account        models.Account
	ExpectedResult bool
}{{
	Description: "Testing a valid account model should validate.",
	Account: models.Account{
		Username: "testusername",
		Email:    "test@email.com",
		Name:     "test name",
		Password: "password1",
	},
	ExpectedResult: true,
}, {
	Description: "Testing account missing an username value should not validate.",
	Account: models.Account{
		Email: "test@email.com",
		Name:  "test name",
	},
	ExpectedResult: false,
}, {
	Description: "Testing account missing an email value should not validate.",
	Account: models.Account{
		Username: "testusername2",
		Name:     "test username",
		Password: "password1",
	},
	ExpectedResult: false,
}, {
	Description: "Testing account missing a name value should not validate.",
	Account: models.Account{
		Username: "testusername2",
		Email:    "email@email.com",
		Password: "password1",
	},
	ExpectedResult: false,
}, {
	Description: "Testing an account missing a password should not validate.",
	Account: models.Account{
		Username: "testusername2",
		Email:    "email@email.com",
		Name:     "test username",
	},
	ExpectedResult: false,
}, {
	Description: "Testing an email in an invalid format should not validate.",
	Account: models.Account{
		Username: "testusername2",
		Email:    "myemail",
		Name:     "test username",
		Password: "password1",
	},
	ExpectedResult: false,
}, {
	Description: "Testing account username length less than 3 should not validate.",
	Account: models.Account{
		Username: "a2",
		Email:    "test@email.com",
		Name:     "test username",
		Password: "password1",
	},
	ExpectedResult: false,
}, {
	Description: "Testing name greater than 32 should not validate.",
	Account: models.Account{
		Username: "testusername3",
		Email:    "test@email.com",
		Name:     "a really long name that is so much longer than 32 characters that the validation is going to fail.",
		Password: "password1",
	},
	ExpectedResult: false,
}, {
	Description: "Testing password less than 8 characters should not validate.",
	Account: models.Account{
		Username: "testusername3",
		Email:    "test@email.com",
		Name:     "my awesome name",
		Password: "2short",
	},
	ExpectedResult: false,
}}

func TestValidateAccount(t *testing.T) {
	t.Log("Testing ValidateAccount...")

	for i, tc := range validateAccountTests {
		t.Logf("[ %02d ] %s\n", i+1, tc.Description)
		a := tc.Account
		err := a.Validate()
		resultValidateMethod := true
		result := tc.ExpectedResult

		if err != nil {
			resultValidateMethod = false
		}

		if result != resultValidateMethod {
			fmt.Println("index: ", i)
			t.Errorf("[ FAIL ] ValidateAccount did not return expected value. Expected \"%v\" but got \"%v\".",
				result, resultValidateMethod)
		}
	}
}

var updateAccountTests = []struct {
	Description            string
	OldAccount, NewAccount models.Account
}{{
	Description: "Testing that updating an account returns the new account.",
	OldAccount: models.Account{
		Username: "OldUsername",
		Email:    "oldemail@email.com",
		Name:     "Old Name",
		Password: "oldpassword",
		Salt:     "oldsalt",
	},
	NewAccount: models.Account{
		Username: "NewUsername",
		Email:    "newemail@email.com",
		Name:     "New Name",
		Password: "oldpassword",
		Salt:     "oldsalt",
	},
}}

func TestUpdateAccount(t *testing.T) {
	t.Log("Testing Updating an Account...")

	for i, tc := range updateAccountTests {
		t.Logf("[ %02d ] %s\n", i+1, tc.Description)
		db.CreateAccount(tc.OldAccount)

		apiErr := UpdateAccount(db, tc.OldAccount.Username, tc.OldAccount.Email, tc.NewAccount)
		if apiErr != nil {
			t.Errorf("[ FAIL ] UpdateAccount failed to update the account. %s\n", apiErr.FullError())
			break
		}

		acc, err := db.GetAccountByUsername(tc.NewAccount.Username)
		if err != nil {
			t.Errorf("[ FAIL ] Failed to get the new account from the database. %s\n", err.Error())
			break
		}

		if !reflect.DeepEqual(tc.NewAccount, acc) {
			t.Errorf("[ FAIL ] The account was not updated correctly.\nExpected %+v\nbut got %+v\n",
				tc.NewAccount, acc)
			break
		}

		db.RemoveAccount(tc.NewAccount.Username, tc.NewAccount.Email)
	}
}
