package diagnostics

import (
	"testing"
)

func TestCalculatesLifeSupportRating(t *testing.T) {

	result := CalculatePowerConsumption(23, 10)
	var expected int64 = 230

	if result != expected {
		t.Error("Did not return expected value. Received", result, "expected ", expected)
	}
}

func TestGetsLifeSupportRating(t *testing.T) {
	values := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}

	result := GetLifeSupportRatingFromDiagnostics(values)
	var expected int64 = 230

	if result != expected {
		t.Error("Did not return expected value. Received", result, "expected ", expected)
	}
}
