package tdd

import (
	"strconv"
)

// StringToInt is function for convert everything string to int
func StringToInt(value string) int {
	valueToInt, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return valueToInt
}
