package navigation

import (
	"strings"
)

const OK = "OK"
const CORRUPTED = "CORRUPTED"
const INCOMPLETE = "INCOMPLETE"

type Reading struct {
	input     string
	status    string
	debugSign string
}

func (reading *Reading) isCorrupted() (bool, string) {
	normalizedSigns := normalizeInput(reading.input)
	signs := strings.Split(normalizedSigns, "")

	startSymbols := []string{}
	for _, sign := range signs {
		if isStartSymbol(sign) {
			startSymbols = append(startSymbols, sign)
			continue
		}

		// Closing symbol
		if !matchesStartSymbol(sign, startSymbols[len(startSymbols)-1]) {
			return true, sign
		}
	}

	return false, ""
}

func (reading *Reading) isIncomplete() (bool, string) {
	normalizedSigns := normalizeInput(reading.input)
	signs := strings.Split(normalizedSigns, "")

	startSymbols := []string{}
	for _, sign := range signs {
		if isStartSymbol(sign) {
			startSymbols = append(startSymbols, sign)
			continue
		}

		if isCloseSymbol(sign) {
			// Remove last match
			startSymbols = removeLastMatchOf(sign, startSymbols)
		}
	}

	return true, strings.Join(startSymbols, "")
}

func (reading *Reading) parseStatus() {
	isCorrupted, debugSign := reading.isCorrupted()
	isInComplete, incompleteSigns := reading.isIncomplete()

	if isCorrupted {
		reading.status = CORRUPTED
		reading.debugSign = debugSign
	} else if isInComplete {
		reading.status = INCOMPLETE
		reading.debugSign = incompleteSigns
	}
}

func GetCorruptedReadings(readings []Reading) []Reading {
	corrupted := []Reading{}

	for _, reading := range readings {
		if reading.status == CORRUPTED {
			corrupted = append(corrupted, reading)
		}
	}

	return corrupted
}

func GetIncompleteReadings(readings []Reading) []Reading {
	incomplete := []Reading{}

	for _, reading := range readings {
		if reading.status == INCOMPLETE {
			incomplete = append(incomplete, reading)
		}
	}

	return incomplete
}

func NewReading(readingValue string) Reading {
	reading := Reading{
		input:  readingValue,
		status: OK,
	}

	reading.parseStatus()

	return reading
}

func NewReadings(lines []string) []Reading {
	readings := []Reading{}

	for _, line := range lines {
		readings = append(readings, NewReading(line))
	}

	return readings
}
