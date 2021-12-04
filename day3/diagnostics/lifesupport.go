package diagnostics

func CalculateLifeSupportRating(oxygen int64, c02 int64) int64 {
	return oxygen * c02
}

func GetLifeSupportRatingFromDiagnostics(values []string) int64 {
	oxygen := GetOxygenRateDecimal(values)
	c02 := GetC02ScrubberRateDecimal(values)

	return CalculateLifeSupportRating(oxygen, c02)
}
