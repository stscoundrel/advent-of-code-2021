package main

import (
	"github.com/stscoundrel/advent-of-code-2021/day13/folds"
	"github.com/stscoundrel/advent-of-code-2021/day13/reader"
)

func CountFirstFold(fileName string) int {
	rawPoints, rawFolds := reader.ReadFoldInstructions(fileName)
	points, instructions := folds.ParseInstructions(rawPoints, rawFolds)
	count := folds.CountSpaces(points, []folds.Fold{instructions[0]})

	return count
}

func DoAllFold(fileName string) int {
	rawPoints, rawFolds := reader.ReadFoldInstructions(fileName)
	points, instructions := folds.ParseInstructions(rawPoints, rawFolds)
	count := folds.CountSpaces(points, instructions)

	return count
}
