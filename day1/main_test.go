package main

import (
	"testing"
)

func TestGetsIncreasesInDepth(t *testing.T) {
	result := GetIncreasesInDepth()
	expected := 1400

	if result != expected {
		t.Error("Did not return expected number. Received", result, "expected ", expected)
	}
}

func TestGetsIncreasesInDepthSliding(t *testing.T) {
	result := GetIncreasesInDepthSliding()
	expected := 1429

	if result != expected {
		t.Error("Did not return expected number. Received", result, "expected ", expected)
	}
}
