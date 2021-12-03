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

func getGammaNumber(frequencies map[int]int) int {
	gammaNumber := 0
	highest := frequencies[0]

	for key, value := range frequencies {
		if value > highest {
			gammaNumber = key
		}
	}

	return gammaNumber
}

func GetGammaRateBinary(values []string) string {
	gammaRate := []string{}
	readings := getReadings(values)
	frequencyMap := map[int]map[int]int{}

	for i := 0; i < 5; i++ {
		frequencyMap[i] = make(map[int]int)
		for _, reading := range readings {
			newValue := 1
			oldValue, hasOldValue := frequencyMap[i][reading[i]]

			if hasOldValue {
				newValue += oldValue
			}

			frequencyMap[i][reading[i]] = newValue
		}

		gammaNumber := getGammaNumber(frequencyMap[i])

		gammaRate = append(gammaRate, strconv.Itoa(gammaNumber))
	}

	return strings.Join(gammaRate[:], "")
}

func GetGammaRateDecimal(values []string) int64 {
	binaryGammaRate := GetGammaRateBinary(values)
	decimal, _ := strconv.ParseInt(binaryGammaRate, 2, 64)

	return decimal
}
