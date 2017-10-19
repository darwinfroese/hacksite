package projects

import (
	"testing"

	"github.com/darwinfroese/hacksite/server/models"
)

var validateProjectTests = []struct {
	Description    string
	Project        models.Project
	ExpectedResult bool
}{{
	Description: "Testing a valid project model should validate.",
	Project: models.Project{
		Name: "Test Project",
	},
	ExpectedResult: true,
}, {
	Description:    "Testing a project missing a name value should not validate.",
	Project:        models.Project{},
	ExpectedResult: false,
}}

func TestValidateProject(t *testing.T) {
	t.Log("Testing ValidateProject...")

	for i, tc := range validateProjectTests {
		t.Logf("[ %02d ] %s\n", i+1, tc.Description)

		result := tc.ValidateProject()
		if result != tc.ExpectedResult {
			t.Errorf("[ FAIL ] ValidateProject did not return expected value. Expected \"%v\" but got \"%v\".",
				result, tc.ExpectedResult)
		}
	}
}
