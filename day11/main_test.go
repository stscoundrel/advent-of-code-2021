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

func TestGetsSimutaneousFlashInTestData(t *testing.T) {
	result := CalculateSimultaneousFlash("./reader/test_input.txt")
	var expected int = 195

	if result != expected {
		t.Error("Did not return expected number. Received", result, "expected ", expected)
	}
}

func TestGetsSimutaneousFlashInData(t *testing.T) {
	result := CalculateSimultaneousFlash("./resources/input.txt")
	var expected int = 403

	if result != expected {
		t.Error("Did not return expected number. Received", result, "expected ", expected)
	}
}
