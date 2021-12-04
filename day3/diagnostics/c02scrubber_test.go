package diagnostics

import (
	"testing"
)

func TestGetsC02RateBinary(t *testing.T) {
	steps := []string{
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

	result := GetC02ScrubberRateBinary(steps)
	expected := "01010"

	if result != expected {
		t.Error("Did not return expected value. Received", result, "expected ", expected)
	}
}

func TestGetsC02RateDecimal(t *testing.T) {
	steps := []string{
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

	result := GetC02ScrubberRateDecimal(steps)
	var expected int64 = 10

	if result != expected {
		t.Error("Did not return expected value. Received", result, "expected ", expected)
	}
}
