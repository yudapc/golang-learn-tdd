package tdd

import (
	"testing"
)

func TestUserStruct(test *testing.T) {
	user := User{Name: "Jhon Khennedy", Age: 35}
	expectedName := "Jhon Khennedy"
	expectedAge := 35
	if user.Name != expectedName || user.Age != expectedAge {
		test.Errorf("The User struct is wrong!")
	}
}
