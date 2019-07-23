package tdd

func IntToUint(value int) uint {
	intFloat := float64(value)
	return uint(intFloat)
}
