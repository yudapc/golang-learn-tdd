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

func TestReduction(test *testing.T) {
	var first = 6
	var second = 3
	var expected = 3
	result := Reduction(first, second)
	if result != expected {
		test.Errorf("The result from function Reduction is wrong!, %d - %d = %d", first, second, expected)
	}
}

func TestDivision(test *testing.T) {
	var first = 9
	var second = 3
	var expected = 3
	result := Division(first, second)
	if result != expected {
		test.Errorf("The result from function Division is wrong!, %d / %d = %d", first, second, expected)
	}
}
