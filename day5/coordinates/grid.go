package coordinates

import (
	"strconv"
)

func getHashForPoint(point Coordinate) string {
	return strconv.Itoa(point.x) + "-" + strconv.Itoa(point.y)
}

func placeLinesOnMap(lines []Line) map[string]int {
	grid := map[string]int{}

	for _, line := range lines {
		if isVerticalOrHorizontal(line) {
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
	}

	return grid
}

func GetOverLappingPoints(lines []Line) (points int) {
	grid := placeLinesOnMap(lines)

	for _, count := range grid {
		if count > 1 {
			points++
		}
	}

	return
}
