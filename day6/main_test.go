package main

import (
	"testing"
)

func TestGetFishCount(t *testing.T) {
	result := CountFishes(80)
	var expected int = 346063

	if result != expected {
		t.Error("Did not return expected number. Received", result, "expected ", expected)
	}
}

func TestGetForeverFishCount(t *testing.T) {
	result := CountFishesInGroups(256)
	var expected int = 1572358335990

	if result != expected {
		t.Error("Did not return expected number. Received", result, "expected ", expected)
	}
}
