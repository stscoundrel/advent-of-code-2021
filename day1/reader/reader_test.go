package reader

import (
	"testing"
)

func TestReadsLinesFromFile(t *testing.T) {
	result := ReadIntsFromFile("./test_input.txt")
	expected := []int{
		5000,
		5001,
		5007,
	}

	for index, _ := range result {
		if result[index] != expected[index] {
			t.Error("Did not return expected content. Received", result[index], "expected ", expected[index])
		}
	}
}
