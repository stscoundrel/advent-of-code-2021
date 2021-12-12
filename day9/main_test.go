package main

import (
	"testing"
)

func TestGetsRisklevel(t *testing.T) {
	result := CalculateRisks("./resources/heightmap.txt")
	var expected int = 486

	if result != expected {
		t.Error("Did not return expected number. Received", result, "expected ", expected)
	}
}

func TestGetsThreelargest(t *testing.T) {
	result := CalculateThreeLargest("./resources/heightmap.txt")
	var expected int = 1059300

	if result != expected {
		t.Error("Did not return expected number. Received", result, "expected ", expected)
	}
}
