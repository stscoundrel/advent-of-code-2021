package signals

import (
	"strconv"
	"strings"
)

func getPresentationNumber(digits string, mapping map[string]int) string {
	return strconv.Itoa(mapping[sortDigits(digits)])
}

func combineOutputNumber(numbers []string) int {
	combinedNumber, _ := strconv.Atoi(strings.Join(numbers, ""))

	return combinedNumber
}

func decipherSignalOutput(signal Signal) int {
	outputDigits := []string{}
	numbersToDigitsMap := getDecipherMap(signal)
	digitsToNumbers := normalizeAndReverse(numbersToDigitsMap)

	for _, digits := range signal.output {
		outputDigits = append(outputDigits, getPresentationNumber(digits, digitsToNumbers))
	}

	return combineOutputNumber(outputDigits)
}

func DecipherSignalOutputs(signals []Signal) int {
	sum := 0

	for _, signal := range signals {
		sum += decipherSignalOutput(signal)
	}

	return sum
}
