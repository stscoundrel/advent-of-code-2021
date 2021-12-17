package main

import (
	"testing"
)

func TestGetsFlashesInTestData1(t *testing.T) {
	result := CalculateFlashes("./reader/test_input.txt", 10)
	var expected int = 204

	if result != expected {
		t.Error("Did not return expected number. Received", result, "expected ", expected)
	}
}

func TestGetsFlashesInTestData2(t *testing.T) {
	result := CalculateFlashes("./reader/test_input.txt", 100)
	var expected int = 1656

	if result != expected {
		t.Error("Did not return expected number. Received", result, "expected ", expected)
	}
}

func TestGetsFlashesInData(t *testing.T) {
	result := CalculateFlashes("./resources/input.txt", 100)
	var expected int = 1688

	if result != expected {
		t.Error("Did not return expected number. Received", result, "expected ", expected)
	}
}
