package folds

import (
	"strconv"
	"strings"
)

func parsePoints(rawPoints [][]int) []Point {
	points := []Point{}

	for _, rawPoint := range rawPoints {
		points = append(points, Point{x: rawPoint[0], y: rawPoint[1]})
	}

	return points
}

func parseFolds(rawFolds []string) []Fold {
	folds := []Fold{}

	for _, rawFold := range rawFolds {
		rawFold = strings.Replace(rawFold, "fold along ", "", -1)
		parts := strings.Split(rawFold, "=")

		value, _ := strconv.Atoi(parts[1])

		folds = append(folds, Fold{direction: parts[0], value: value})
	}

	return folds
}

func ParseInstructions(points [][]int, folds []string) ([]Point, []Fold) {
	return parsePoints(points), parseFolds(folds)
}
