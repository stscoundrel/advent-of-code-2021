package signals

func isUniqueDigit(digit string) bool {
	length := len(digit)

	if length == 7 || length == 4 || length == 3 || length == 2 {
		return true
	}

	return false
}

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
