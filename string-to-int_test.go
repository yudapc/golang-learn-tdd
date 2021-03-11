package tdd

import (
	"testing"
)

func TestStringToInt(t *testing.T) {
	t.Run("should return int when string is number", func(t *testing.T) {
		input := "12"
		value := StringToInt(input)
		if value != 12 {
			t.Errorf("Value type is not int")
		}
	})
	t.Run("should return 0 when invalid number", func(t *testing.T) {
		input := "ngaco"
		value := StringToInt(input)
		if value != 0 {
			t.Errorf("Value type is not int")
		}
	})
	t.Run("should return 0 when combination string and number", func(t *testing.T) {
		input := "ngaco1"
		value := StringToInt(input)
		if value != 0 {
			t.Errorf("Value type is not int")
		}
	})
}
