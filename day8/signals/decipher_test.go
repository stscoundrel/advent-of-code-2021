package signals

import (
	"testing"
)

func TestDeciphersOutputFromSignal(t *testing.T) {
	signal := Signal{
		input:  []string{"acedgfb", "cdfbe", "gcdfa", "fbcad", "dab", "cefabd", "cdfgeb", "eafb", "cagedb", "ab"},
		output: []string{"cdfeb", "fcadb", "cdfeb", "cdbaf"},
	}

	result := decipherSignalOutput(signal)
	expected := 5353

	if result != expected {
		t.Error("Did not return expected content. Received", result, "expected ", expected)
	}

}
