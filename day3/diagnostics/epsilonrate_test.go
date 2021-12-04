package diagnostics

import (
	"testing"
)

func TestGetsEpsilonRateBinary(t *testing.T) {
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

	result := GetEpsilonRateBinary(steps)
	expected := "01001"

	if result != expected {
		t.Error("Did not return expected value. Received", result, "expected ", expected)
	}
}

func TestGetsEpsilonRateDecimal(t *testing.T) {
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

	result := GetEpsilonRateDecimal(steps)
	var expected int64 = 9

	if result != expected {
		t.Error("Did not return expected value. Received", result, "expected ", expected)
	}
}
