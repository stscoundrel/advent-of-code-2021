package navigation

import (
	"sort"
	"strings"
)

var errorScoreMap map[string]int = map[string]int{
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
}

var missingScoreMap map[string]int = map[string]int{
	")": 1,
	"]": 2,
	"}": 3,
	">": 4,
}

func ScoreSyntaxErrors(readings []Reading) int {
	score := 0

	for _, reading := range readings {
		if reading.status == CORRUPTED {
			score += errorScoreMap[reading.debugSign]
		}
	}

	return score
}

func scoreMissingEnding(reading Reading) int {
	score := 0

	signs := strings.Split(reading.debugSign, "")
	reverseOpens(signs)

	for _, missingSign := range signs {
		score *= 5
		score += missingScoreMap[openSymbols[missingSign]]
	}

	return score
}

func ScoreMissingEndings(readings []Reading) []int {
	scores := []int{}

	for _, reading := range readings {
		if reading.status == INCOMPLETE {
			scores = append(scores, scoreMissingEnding(reading))
		}
	}

	return scores
}

func GetMiddleScore(scores []int) int {
	sort.Slice(scores, func(i, j int) bool {
		return scores[i] < scores[j]
	})

	return scores[len(scores)/2]
}
