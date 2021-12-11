package signals

import (
	"testing"
)

func TestTransformsLinesToSignals(t *testing.T) {
	lines := []string{
		"be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe",
	}

	result := LinesToSignals(lines)

	expectedInput := []string{
		"be", "cfbegad", "cbdgef", "fgaecd", "cgeb", "fdcge", "agebfd", "fecdb", "fabcd", "edb",
	}

	expectedOutput := []string{
		"fdgacbe", "cefdb", "cefbgd", "gcbe",
	}

	for _, signal := range result {
		for index, input := range signal.input {
			if input != expectedInput[index] {
				t.Error("Did not return expected content. Received", input, "expected ", expectedInput[index])
			}
		}

		for index, output := range signal.output {
			if output != expectedOutput[index] {
				t.Error("Did not return expected content. Received", output, "expected ", expectedOutput[index])
			}
		}
	}

}
