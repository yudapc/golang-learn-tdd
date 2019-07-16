package tdd

import (
	"testing"
)

func TestAddition(test *testing.T) {
	var first = 5
	var second = 6
	var expected = 11
	result := Addition(first, second)
	if result != expected {
		test.Errorf("The result from function Additional is wrong!, %d + %d = %d", first, second, expected)
	}
}
