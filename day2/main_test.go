package main

import (
	"testing"
)

func TestLocatesShip(t *testing.T) {
	result := LocateShip()
	expected := 1250395

	if result != expected {
		t.Error("Did not return expected number. Received", result, "expected ", expected)
	}
}
