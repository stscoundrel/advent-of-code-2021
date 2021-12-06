package reader

import (
	"testing"
)

func TestReadsLinesFromFile(t *testing.T) {
	result := ReadLinesFromFile("./test_input.txt")
	expected := []string{
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

	for index, _ := range result {
		if result[index] != expected[index] {
			t.Error("Did not return expected content. Received", result[index], "expected ", expected[index])
		}
	}
}
