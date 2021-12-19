package main

import (
	"testing"
)

func TestCountsFirstFoldInTestInput(t *testing.T) {
	result := CountFirstFold("./reader/test_input.txt")
	var expected int = 17

	if result != expected {
		t.Error("Did not return expected number. Received", result, "expected ", expected)
	}
}

func TestCountsFirstFoldInInput(t *testing.T) {
	result := CountFirstFold("./resources/input.txt")
	var expected int = 675

	if result != expected {
		t.Error("Did not return expected number. Received", result, "expected ", expected)
	}
}

func TestDoesALlTheFolds(t *testing.T) {
	result := DoAllFold("./resources/input.txt")
	var expected int = 98

	// Dots are irrelevant, outputted form matters: It's:

	// [1 0 0 1 0 1 1 1 1 0 1 0 0 1 0 1 0 0 1 0 1 1 1 1 0 1 1 1 1 0 0 0 1 1 0 1 1 1 1]
	// [1 0 0 1 0 0 0 0 1 0 1 0 1 0 0 1 0 0 1 0 1 0 0 0 0 1 0 0 0 0 0 0 0 1 0 0 0 0 1]
	// [1 1 1 1 0 0 0 1 0 0 1 1 0 0 0 1 1 1 1 0 1 1 1 0 0 1 1 1 0 0 0 0 0 1 0 0 0 1 0]
	// [1 0 0 1 0 0 1 0 0 0 1 0 1 0 0 1 0 0 1 0 1 0 0 0 0 1 0 0 0 0 0 0 0 1 0 0 1 0 0]
	// [1 0 0 1 0 1 0 0 0 0 1 0 1 0 0 1 0 0 1 0 1 0 0 0 0 1 0 0 0 0 1 0 0 1 0 1 0 0 0]
	// [1 0 0 1 0 1 1 1 1 0 1 0 0 1 0 1 0 0 1 0 1 0 0 0 0 1 1 1 1 0 0 1 1 0 0 1 1 1 1]

	// Highlighting ones gives H Z K H F E J Z

	if result != expected {
		t.Error("Did not return expected number. Received", result, "expected ", expected)
	}
}
