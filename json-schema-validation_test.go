package tdd

import "testing"

func TestValidationUsingJSONSchema(test *testing.T) {
	var jsonString = `{"email": "jhon@doe.com", "password": "securewords"}`
	var expected = "The document is valid"
	result, _ := ValidationUsingJSONSchema(jsonString)
	if result != expected {
		test.Errorf("The result from ValidationUsingJSONSchema is wrong!")
	}
}

func TestValidationUsingJSONSchemaHasError(test *testing.T) {
	var jsonString = `{"email": "jhon@doe.com", "passwords": "securewords"}`
	_, err := ValidationUsingJSONSchema(jsonString)
	if err == nil {
		test.Errorf("The result from ValidationUsingJSONSchema is not error!")
	}
}

func TestValidationUsingJSONSchemaHasErrorInput(test *testing.T) {
	var jsonString = `{"email": "jhon@doe.com", "passwords": "securewords"`
	_, err := ValidationUsingJSONSchema(jsonString)
	if err == nil {
		test.Errorf("The result from ValidationUsingJSONSchema is not error!")
	}
}
