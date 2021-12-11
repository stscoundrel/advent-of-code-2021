package signals

import (
	"strings"
)

func getInputAndOutput(line string) (input []string, output []string) {
	parts := strings.Split(line, " | ")

	input = strings.Split(parts[0], " ")
	output = strings.Split(parts[1], " ")

	return input, output
}

func LinesToSignals(lines []string) []Signal {
	signals := []Signal{}

	for _, line := range lines {
		input, output := getInputAndOutput(line)

		signals = append(signals, Signal{input, output})
	}

	return signals
}
