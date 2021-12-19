package main

import (
	"testing"
)

func TestGetsMostAndLeastCommonDifferenceInTestInput(t *testing.T) {
	result := GetMostAndLeastCommonDifference("./reader/test_input.txt", 10)
	var expected int = 1588

	if result != expected {
		t.Error("Did not return expected number. Received", result, "expected ", expected)
	}
}

func TestGetsMostAndLeastCommonDifferenceInInput(t *testing.T) {
	result := GetMostAndLeastCommonDifference("./resources/input.txt", 10)
	var expected int = 3259

	if result != expected {
		t.Error("Did not return expected number. Received", result, "expected ", expected)
	}
}
