package tdd

import (
	"reflect"
	"testing"
)

func TestStringToInt(test *testing.T) {
	value := "12"
	valueType := StringToInt(value)
	typeValue := reflect.TypeOf(valueType).Kind()
	if typeValue != reflect.Int {
		test.Errorf("Value type is not int")
	}
}
