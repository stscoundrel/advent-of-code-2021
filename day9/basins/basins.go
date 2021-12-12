package basins

import (
	"sort"

	"github.com/stscoundrel/advent-of-code-2021/day9/heights"
)

func GetInitialBasins(heightMap [][]int) []Basin {
	basins := []Basin{}

	for x := 0; x < len(heightMap); x += 1 {
		for y := 0; y < len(heightMap[x]); y += 1 {
			if heights.IsLowPoint(x, y, heightMap) {
				lowPoint := NewPoint(x, y, heightMap[x][y])
				basins = append(basins, Basin{points: []Point{lowPoint}})
			}
		}
	}

	return basins
}

func getPointNeighbours(point Point, basin Basin, heightMap [][]int) []Point {
	newPoints := []Point{}

	if point.canExpandUp(heightMap, basin) {
		newPoints = append(newPoints, NewPointUp(point, heightMap))
	}

	if point.canExpandDown(heightMap, basin) {
		newPoints = append(newPoints, NewPointDown(point, heightMap))
	}

	if point.canExpandLeft(heightMap, basin) {
		newPoints = append(newPoints, NewPointLeft(point, heightMap))
	}

	if point.canExpandRight(heightMap, basin) {
		newPoints = append(newPoints, NewPointRight(point, heightMap))
	}

	return newPoints
}

func FillABasin(basin Basin, heightMap [][]int) Basin {
	basinIsNotFilled := true

	for basinIsNotFilled {
		if basin.isFull() {
			basinIsNotFilled = false
		}

		for index, point := range basin.points {
			if point.shouldExpand() {
				neighbours := getPointNeighbours(point, basin, heightMap)
				basin.points[index].hasExpanded = true
				basin.points = append(basin.points, neighbours...)
			}
		}
	}

	return basin
}

func FillBasins(basins []Basin, heightMap [][]int) []Basin {
	filledBasins := []Basin{}

	for _, basin := range basins {
		filledBasins = append(filledBasins, FillABasin(basin, heightMap))
	}

	return filledBasins
}

func GetBasins(heightMap [][]int) []Basin {
	initialBasins := GetInitialBasins(heightMap)
	basins := FillBasins(initialBasins, heightMap)

	return basins
}

func getThreeLargest(basins []Basin) []Basin {
	sort.Slice(basins, func(i, j int) bool {
		return basins[i].getSize() > basins[j].getSize()
	})

	return []Basin{basins[0], basins[1], basins[2]}
}

func MultiplyThreeLargest(basins []Basin) int {
	three := getThreeLargest(basins)

	return three[0].getSize() * three[1].getSize() * three[2].getSize()
}
