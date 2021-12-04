package diagnostics

func getEpsilonNumber(frequencies map[int]int) int {
	epsilonNumber := 0
	highest := frequencies[0]

	for key, value := range frequencies {
		if value < highest {
			epsilonNumber = key
		}
	}

	return epsilonNumber
}

func GetEpsilonRateBinary(values []string) string {
	return getRateBinary(values, getEpsilonNumber)
}

func GetEpsilonRateDecimal(values []string) int64 {
	binaryEpsilonRate := GetEpsilonRateBinary(values)
	return BinaryRateToDecimal(binaryEpsilonRate)
}
