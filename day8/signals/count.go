package signals

func CountUniquesInOutputs(signals []Signal) int {
	count := 0

	for _, signal := range signals {
		for _, outputDigit := range signal.output {
			if isUniqueDigit(outputDigit) {
				count++
			}
		}
	}

	return count
}
