package main

import (
	"testing"
)

func TestPlaysWithTheSquid(t *testing.T) {
	result := PlaWithTheSquid()
	var expected int = 28082

	if result != expected {
		t.Error("Did not return expected number. Received", result, "expected ", expected)
	}
}

func TestLetsTheSquidWin(t *testing.T) {
	result := LetTheSquidWin()
	var expected int = 8224

	if result != expected {
		t.Error("Did not return expected number. Received", result, "expected ", expected)
	}
}
