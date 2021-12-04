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

func getRateBinary(values []string, rateGetter func(map[int]int) int) string {
	rate := []string{}
	readings := getReadings(values)
	frequencyMap := map[int]map[int]int{}

	for i := 0; i < len(values[0]); i++ {
		frequencyMap[i] = make(map[int]int)
		for _, reading := range readings {
			newValue := 1
			oldValue, hasOldValue := frequencyMap[i][reading[i]]

			if hasOldValue {
				newValue += oldValue
			}

			frequencyMap[i][reading[i]] = newValue
		}

		rateNumber := rateGetter(frequencyMap[i])

		rate = append(rate, strconv.Itoa(rateNumber))
	}

	return strings.Join(rate[:], "")
}

func BinaryRateToDecimal(binaryRate string) int64 {
	decimal, _ := strconv.ParseInt(binaryRate, 2, 64)

	return decimal
}
