package reader

import (
	"testing"
)

func TestReadsLinesFromFile(t *testing.T) {
	result := ReadStringsFromFile("./test_input.txt")
	expected := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}

	for index, _ := range result {
		if result[index] != expected[index] {
			t.Error("Did not return expected content. Received", result[index], "expected ", expected[index])
		}
	}
}
