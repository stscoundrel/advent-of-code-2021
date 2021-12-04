package diagnostics

func CalculatePowerConsumption(gammaRate int64, epsilonRate int64) int64 {
	return gammaRate * epsilonRate
}

func GetPowerConsumptionFromDiagnostics(values []string) int64 {
	gammaRate := GetGammaRateDecimal(values)
	epsilonRate := GetEpsilonRateDecimal(values)

	return CalculatePowerConsumption(gammaRate, epsilonRate)
}
