package tdd

import "strconv"

// StringToUint function
func StringToUint(value string) uint {
	toInt, _ := strconv.Atoi(value)
	toFloat := float64(toInt)
	toUint := uint(toFloat)
	return toUint
}
