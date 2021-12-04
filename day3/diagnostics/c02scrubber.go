package diagnostics

func GetC02ScrubberRateBinary(values []string) string {
	return getFilteredRateBy(values, getLeastCommonInPosition)
}

func GetC02ScrubberRateDecimal(values []string) int64 {
	binaryOxygenRate := GetC02ScrubberRateBinary(values)
	return BinaryRateToDecimal(binaryOxygenRate)
}
