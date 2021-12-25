package main

import (
	"testing"
)

func TestGetsMostEfficientPathInTestInput(t *testing.T) {
	result := GetMostEfficientPath("./reader/test_input.txt")
	var expected int = 40

	if result != expected {
		t.Error("Did not return expected number. Received", result, "expected ", expected)
	}
}

func TestGetsMostEfficientPathInTestInput2(t *testing.T) {
	result := GetMostEfficientPath("./reader/test_input2.txt")
	var expected int = 315

	if result != expected {
		t.Error("Did not return expected number. Received", result, "expected ", expected)
	}
}

func TestGetsMostEfficientPathInInput(t *testing.T) {
	result := GetMostEfficientPath("./resources/input.txt")
	var expected int = 609

	if result != expected {
		t.Error("Did not return expected number. Received", result, "expected ", expected)
	}
}

func TestGetsMostEfficientPathInFullInput(t *testing.T) {
	result := GetMostEfficientFullPath("./resources/input.txt")
	var expected int = 2925

	if result != expected {
		t.Error("Did not return expected number. Received", result, "expected ", expected)
	}
}
