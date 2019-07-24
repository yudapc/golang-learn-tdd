package tdd

import (
	"testing"
)

func TestStringToSnakeCase(test *testing.T) {
	camelCase := "ThisIsCamelCaseName"
	expected := "this_is_camel_case_name"
	result := StringToSnakeCase(camelCase)
	if result != expected {
		test.Error("Wrong convert to snake case")
	}
}
