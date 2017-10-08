package auth

import (
	"fmt"
	"testing"
)

var loginTests = []struct {
	Description        string
	Username, Password string
	ExpectedError      string
	ExpectedID         int
}{{
	Description:   "Logging in with valid account should succeed.",
	Username:      username,
	Password:      password,
	ExpectedError: "",
	ExpectedID:    id,
}, {
	Description:   "Logging in with an invalid account should not succeed.",
	Username:      username,
	Password:      "bad-password",
	ExpectedError: fmt.Sprintf("Error: %s", UnathorizedErrorMessage),
}}

func TestLogin(t *testing.T) {
	t.Log("Testing Login...")

	for i, tc := range loginTests {
		t.Logf("[ %02d ] %s\n", i+1, tc.Description)

		sesh, err := Login(db, tc.Username, tc.Password)

		if err != nil && err.Error() != tc.ExpectedError {
			t.Errorf("[ FAIL ] There was an unexpected error logging in: %s\n", err.Error())
			break
		}

		if tc.ExpectedError == "" {
			if tc.ExpectedID != sesh.UserID {
				t.Errorf("[ FAIL ] The wrong user id was returned. Expected %d but got %d\n", tc.ExpectedID, sesh.UserID)
			}
		}
	}
}
