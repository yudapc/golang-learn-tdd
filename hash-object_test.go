package tdd

import (
	"testing"
)

func TestHashObject(test *testing.T) {
	value := HashObject{
		"ID":   1,
		"Name": "Hash Name",
	}
	if value["ID"] != 1 || value["Name"] != "Hash Name" {
		test.Error("Wrong convert to hash object!")
	}
}
