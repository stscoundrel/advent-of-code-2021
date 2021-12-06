package coordinates

import (
	"testing"
)

func TestReadsLinesFromFile(t *testing.T) {
	rawLines := []string{
		"0,9 -> 5,9",
		"8,0 -> 0,8",
		"9,4 -> 3,4",
		"2,2 -> 2,1",
		"7,0 -> 7,4",
		"6,4 -> 2,0",
		"0,9 -> 2,9",
		"3,4 -> 1,4",
		"0,0 -> 8,8",
		"5,5 -> 8,2",
	}

	result := ParseCoordinates(rawLines)

	if result[0].start.x != 0 {
		t.Error("Did not return expected content. Received", result[0].start.x, "expected ", 0)
	}

	if result[0].start.y != 9 {
		t.Error("Did not return expected content. Received", result[0].start.y, "expected ", 9)
	}

	if result[0].end.x != 5 {
		t.Error("Did not return expected content. Received", result[0].end.x, "expected ", 5)
	}

	if result[0].end.y != 9 {
		t.Error("Did not return expected content. Received", result[0].end.y, "expected ", 9)
	}

}
