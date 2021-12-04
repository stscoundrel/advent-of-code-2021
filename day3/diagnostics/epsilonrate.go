package diagnostics

func GetEpsilonRateBinary(values []string) string {
	return getRateBinary(values, getLeastCommonInPosition)
}

func GetEpsilonRateDecimal(values []string) int64 {
	binaryEpsilonRate := GetEpsilonRateBinary(values)
	return BinaryRateToDecimal(binaryEpsilonRate)
}
