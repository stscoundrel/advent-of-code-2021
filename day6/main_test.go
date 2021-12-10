package main

import (
	"testing"
)

func TestGetFishCount(t *testing.T) {
	result := CountFishes()
	var expected int = 346063

	if result != expected {
		t.Error("Did not return expected number. Received", result, "expected ", expected)
	}
}
