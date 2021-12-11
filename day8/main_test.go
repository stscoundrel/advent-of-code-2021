package main

import (
	"testing"
)

func TestCountsUniquesInTestOutout(t *testing.T) {
	result := CountUniquesInOutput("./reader/test_input.txt")
	var expected int = 26

	if result != expected {
		t.Error("Did not return expected number. Received", result, "expected ", expected)
	}
}

func TestCountsUniquesInOutout(t *testing.T) {
	result := CountUniquesInOutput("./resources/signals.txt")
	var expected int = 539

	if result != expected {
		t.Error("Did not return expected number. Received", result, "expected ", expected)
	}
}
