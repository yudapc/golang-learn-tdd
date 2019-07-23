package tdd

import (
	"strconv"
)

func StringToInt(value string) int {
	valueToInt, _ := strconv.Atoi(value)
	return valueToInt
}
