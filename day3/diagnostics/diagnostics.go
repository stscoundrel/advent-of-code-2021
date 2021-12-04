package diagnostics

import (
	"strconv"
	"strings"
)

func getReadings(values []string) (readings [][]int) {
	for _, value := range values {
		var reading []int
		for _, letterByte := range value {
			letter, _ := strconv.Atoi(string(letterByte))
			reading = append(reading, letter)
		}

		readings = append(readings, reading)
	}

	return
}

func getMostOrLeastCommonInPosition(values [][]int, position int, most bool) int {
	zerosCount := 0
	onesCount := 0

	for _, binarynumber := range values {
		if binarynumber[position] == 0 {
			zerosCount++
		} else {
			onesCount++
		}
	}
	if most {
		if zerosCount > onesCount {
			return 0
		} else {
			return 1
		}
	}

	if zerosCount <= onesCount {
		return 0
	}

	return 1

}

func getMostCommonInPosition(values [][]int, position int) int {
	return getMostOrLeastCommonInPosition(values, position, true)
}

func getLeastCommonInPosition(values [][]int, position int) int {
	return getMostOrLeastCommonInPosition(values, position, false)
}

func getRateBinary(values []string, rateGetter func([][]int, int) int) string {
	rate := []string{}
	readings := getReadings(values)

	for i := 0; i < len(values[0]); i++ {
		intRate := rateGetter(readings, i)
		strRate := strconv.Itoa(intRate)
		rate = append(rate, strRate)
	}

	return strings.Join(rate[:], "")
}

func filter(values [][]int, compareGetter func([][]int, int) int, position int) []int {
	if len(values) == 1 {
		return values[0]
	}

	var filtered = [][]int{}
	comparator := compareGetter(values, position)
	for _, line := range values {
		if line[position] == comparator {
			filtered = append(filtered, line)
		}
	}

	return filter(filtered, compareGetter, position+1)
}

func getFilteredRateBy(values []string, compareGetter func([][]int, int) int) string {
	readings := getReadings(values)
	filtered := filter(readings, compareGetter, 0)
	rate := ""

	for _, number := range filtered {
		rate = rate + strconv.Itoa(number)
	}

	return rate
}

func BinaryRateToDecimal(binaryRate string) int64 {
	decimal, _ := strconv.ParseInt(binaryRate, 2, 64)

	return decimal
}
