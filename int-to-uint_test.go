package tdd

import (
	"reflect"
	"testing"
)

func TestIntToUint(test *testing.T) {
	var value int = 42
	valueType := IntToUint(value)
	typeValue := reflect.TypeOf(valueType).Kind()
	if typeValue != reflect.Uint {
		test.Errorf("Value type is not uint")
	}
}
