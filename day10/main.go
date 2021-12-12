package main

import (
	"github.com/stscoundrel/advent-of-code-2021/day10/navigation"
	"github.com/stscoundrel/advent-of-code-2021/day10/reader"
)

func ScoreCorruptionErrors(fileName string) int {
	lines := reader.ReadNavigationInputFromFile(fileName)
	readings := navigation.NewReadings(lines)
	score := navigation.ScoreSyntaxErrors(readings)

	return score
}

func ScoreMissingEnds(fileName string) int {
	lines := reader.ReadNavigationInputFromFile(fileName)
	readings := navigation.NewReadings(lines)
	incompletes := navigation.GetIncompleteReadings(readings)
	scores := navigation.ScoreMissingEndings(incompletes)
	score := navigation.GetMiddleScore(scores)

	return score
}
