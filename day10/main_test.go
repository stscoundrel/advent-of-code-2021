package main

import "testing"

func TestScoresCorruptedInTestInput(t *testing.T) {
	result := ScoreCorruptionErrors("./reader/test_input.txt")
	var expected int = 26397

	if result != expected {
		t.Error("Did not return expected number. Received", result, "expected ", expected)
	}
}

func TestScoresCorruptedInInput(t *testing.T) {
	result := ScoreCorruptionErrors("./resources/input.txt")
	var expected int = 339537

	if result != expected {
		t.Error("Did not return expected number. Received", result, "expected ", expected)
	}
}
