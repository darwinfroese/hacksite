package auth

import (
	"fmt"
	"testing"
)

var loginTests = []struct {
	Description        string
	Username, Password string
	ExpectedError      string
}{{
	Description:   "Logging in with valid account should succeed.",
	Username:      username,
	Password:      password,
	ExpectedError: "",
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
			if tc.Username != sesh.Username {
				t.Errorf("[ FAIL ] The wrong user username was returned. Expected %s but got %s\n", tc.Username, sesh.Username)
			}
		}
	}
}
