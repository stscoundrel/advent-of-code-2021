package locator

import (
	"testing"
)

func TestLocatesPosition(t *testing.T) {
	steps := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}

	result := Locate(steps)

	if result != 150 {
		t.Error("Did not return expected value. Received", result, "expected 7")
	}

}
