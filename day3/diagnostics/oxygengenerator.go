package diagnostics

func GetOxygenRateBinary(values []string) string {
	return getFilteredRateBy(values, getMostCommonInPosition)
}

func GetOxygenRateDecimal(values []string) int64 {
	binaryOxygenRate := GetOxygenRateBinary(values)
	return BinaryRateToDecimal(binaryOxygenRate)
}
