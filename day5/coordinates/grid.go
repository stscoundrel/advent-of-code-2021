package coordinates

import (
	"strconv"
)

func getHashForPoint(point Coordinate) string {
	return strconv.Itoa(point.x) + "-" + strconv.Itoa(point.y)
}

func filterLines(lines []Line, diagonal bool) []Line {
	filteredLines := []Line{}

	for _, line := range lines {
		if diagonal {
			filteredLines = append(filteredLines, line)
		} else {
			if isVerticalOrHorizontal(line) {
				filteredLines = append(filteredLines, line)
			}
		}
	}

	return filteredLines
}

func placeLinesOnMap(lines []Line, diagonal bool) map[string]int {
	grid := map[string]int{}
	filteredLines := filterLines(lines, diagonal)

	for _, line := range filteredLines {
		linePoints := GetPointsOfLine(line)

		for _, linePoint := range linePoints {
			pointHash := getHashForPoint(linePoint)
			count, hasCoordinate := grid[pointHash]

			if hasCoordinate {
				grid[pointHash] = count + 1
			} else {
				grid[pointHash] = 1
			}

		}
	}

	return grid
}

func GetOverLappingPoints(lines []Line, diagonal bool) (points int) {
	grid := placeLinesOnMap(lines, diagonal)

	for _, count := range grid {
		if count > 1 {
			points++
		}
	}

	return
}
