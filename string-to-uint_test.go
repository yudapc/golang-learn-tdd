package tdd

import (
	"reflect"
	"testing"
)

func TestStringToUint(test *testing.T) {
	value := "12"
	valueType := StringToUint(value)
	typeValue := reflect.TypeOf(valueType).Kind()
	if typeValue != reflect.Uint {
		test.Errorf("Value type is not uint")
	}
}
