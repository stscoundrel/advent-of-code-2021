package main

import (
	"testing"
)

func TestCountsPathsInTestInput(t *testing.T) {
	result := CountPaths("./reader/test_input.txt", 10)
	var expected int = 10

	if result != expected {
		t.Error("Did not return expected number. Received", result, "expected ", expected)
	}
}

func TestCountsPathsInInput(t *testing.T) {
	result := CountPaths("./resources/input.txt", 10)
	var expected int = 10

	if result != expected {
		t.Error("Did not return expected number. Received", result, "expected ", expected)
	}
}
