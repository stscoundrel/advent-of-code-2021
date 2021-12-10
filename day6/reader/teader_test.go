package reader

import (
	"testing"
)

func TestReadsLinesFromFile(t *testing.T) {
	result := ReadIntsFromFile("./test_input.txt")
	expected := []int{
		3, 4, 3, 1, 2,
	}

	for index, _ := range result {
		if result[index] != expected[index] {
			t.Error("Did not return expected content. Received", result[index], "expected ", expected[index])
		}
	}
}
