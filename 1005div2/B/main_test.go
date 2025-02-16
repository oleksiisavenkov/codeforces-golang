package main

import (
	"testing"
)

func TestCalcUniqueElements(t *testing.T) {

	testCase := func(input []int, expected int) {
		uniqueElements := calcUniqueElements(input)
		if len(uniqueElements) != expected {
			t.Errorf("Expected %d unique elements, got %d", expected, len(uniqueElements))
		}
	}

	testCase([]int{1, 2, 3, 4, 5}, 5)
	testCase([]int{1, 1, 1, 1, 1}, 1)
	testCase([]int{1, 2, 1, 2, 1}, 2)
	testCase([]int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}, 5)
	testCase([]int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5, 1, 2, 3, 4, 5}, 5)
	testCase([]int{}, 0)
}

func TestCalcMaxConsecutiveElementsRange(t *testing.T) {

	testCase := func(input []int, expectedStart int, expectedEnd int, expected bool) {
		start, end, ok := calcMaxConsecutiveElementsRange(input)
		if start != expectedStart || end != expectedEnd || ok != expected {
			t.Errorf("Expected (%d, %d, %t), got (%d, %d, %t)", expectedStart, expectedEnd, expected, start, end, ok)
		}
	}

	testCase([]int{1, 2, 3, 4, 5}, 0, 4, true)
	testCase([]int{1, 1, 1, 1, 1}, 0, 0, false)
	testCase([]int{1, 2, 3, 3, 3}, 0, 1, true)
}

func TestSolve(t *testing.T) {

	testCase := func(input string, expected string) {
		result := solve(input)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	}

	testCase("1 2 3 4 5", "1 5")
	testCase("1 1 1 1 1", "0")
	testCase("1 2 3 3 3", "1 2")
	testCase("1", "1 1")
	testCase("2 1 3 2", "2 3")
}
