package diagnostics

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
	return getRateBinary(values, getGammaNumber)
}

func GetGammaRateDecimal(values []string) int64 {
	binaryGammaRate := GetGammaRateBinary(values)
	return BinaryRateToDecimal(binaryGammaRate)
}
