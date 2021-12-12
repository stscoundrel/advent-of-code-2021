package navigation

var scoreMap map[string]int = map[string]int{
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
}

func ScoreSyntaxErrors(readings []Reading) int {
	score := 0

	for _, reading := range readings {
		if reading.status == CORRUPTED {
			score += scoreMap[reading.errorSign]
		}
	}

	return score
}
