package signals

import (
	"strings"
)

func isUniqueDigit(digit string) bool {
	length := len(digit)

	if length == 7 || length == 4 || length == 3 || length == 2 {
		return true
	}

	return false
}

func getUniqueDigit(digit string) int {
	length := len(digit)

	if length == 7 {
		return 8
	}

	if length == 4 {
		return 4
	}

	if length == 3 {
		return 7
	}

	return 1
}

func getInitialMap(signal Signal) map[int]string {
	mapping := map[int]string{}

	for _, digit := range signal.input {
		if isUniqueDigit(digit) {
			mapping[getUniqueDigit(digit)] = digit
		}
	}

	return mapping
}

func isThree(digits string, oneDigit string) bool {
	containsOne := true

	for _, digit := range oneDigit {
		if !strings.Contains(digits, string(digit)) {
			containsOne = false
		}
	}

	return containsOne
}

func isZero(digits string, oneDigit string) bool {
	containsOne := true

	for _, digit := range oneDigit {
		if !strings.Contains(digits, string(digit)) {
			containsOne = false
		}
	}

	return containsOne
}

func isNine(digits string, fourDigit string) bool {
	containsFour := true

	for _, digit := range fourDigit {
		if !strings.Contains(digits, string(digit)) {
			containsFour = false
		}
	}

	return containsFour
}

func isFive(digits string, fourDigit string, oneDigit string) bool {
	missedFourDigits := []string{}
	containsOne := true

	for _, digit := range oneDigit {
		if !strings.Contains(digits, string(digit)) {
			containsOne = false
		}
	}

	for _, digit := range fourDigit {
		if !strings.Contains(digits, string(digit)) {
			missedFourDigits = append(missedFourDigits, string(digit))
		}
	}

	return len(missedFourDigits) == 1 && !containsOne
}

func getDecipherMap(signal Signal) map[int]string {
	mapping := getInitialMap((signal))

	for _, digit := range signal.input {
		if !isUniqueDigit(digit) {
			length := len(digit)

			if length == 6 {
				if isNine(digit, mapping[4]) {
					mapping[9] = digit
					continue
				}

				if isZero(digit, mapping[1]) {
					mapping[0] = digit
					continue
				}

				mapping[6] = digit
				continue
			}

			if length == 5 {
				if isThree(digit, mapping[1]) {
					mapping[3] = digit
					continue
				}

				if isFive(digit, mapping[4], mapping[1]) {
					mapping[5] = digit
					continue
				}

				mapping[2] = digit
				continue
			}

		}
	}

	return mapping
}
