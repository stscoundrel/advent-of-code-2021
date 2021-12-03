package reader

import (
	"testing"
)

func TestReadsLinesFromFile(t *testing.T) {
	result := ReadFile("./test_input.txt")
	expected := []string{
		"5000",
		"5001",
		"5007",
	}

	for index, _ := range result {
		if result[index] != expected[index] {
			t.Error("Did not return expected content. Received", result[index], "expected ", expected[index])
		}
	}
}
