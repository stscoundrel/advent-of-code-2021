package navigation

import (
	"strings"
)

const OK = "OK"
const CORRUPTED = "CORRUPTED"
const INCOMPLETE = "INCOMPLETE"
const UNKOWN = "UNKNOWN"

type Reading struct {
	input     string
	status    string
	errorSign string
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

func (reading *Reading) parseStatus() {
	isCorrupted, errorSign := reading.isCorrupted()

	if isCorrupted {
		reading.status = CORRUPTED
		reading.errorSign = errorSign
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
