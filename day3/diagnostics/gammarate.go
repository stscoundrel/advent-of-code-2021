package diagnostics

func GetGammaRateBinary(values []string) string {
	return getRateBinary(values, getMostCommonInPosition)
}

func GetGammaRateDecimal(values []string) int64 {
	binaryGammaRate := GetGammaRateBinary(values)
	return BinaryRateToDecimal(binaryGammaRate)
}
