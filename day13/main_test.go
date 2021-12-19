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
