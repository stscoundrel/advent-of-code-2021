package folds

import "fmt"

func getSize(points []Point) (int, int) {
	width := points[0].x
	height := points[0].y

	for _, point := range points {
		if point.x > width {
			width = point.x
		}
		if point.y > height {
			height = point.y
		}
	}

	return width + 1, height + 1
}

func getPointGrid(points []Point) [][]int {
	width, height := getSize(points)
	grid := [][]int{}

	for i := 0; i < height; i += 1 {
		grid = append(grid, make([]int, width))
	}

	// for _, row := range grid {
	// 	fmt.Println(row)
	// }

	for _, point := range points {
		grid[point.y][point.x] = 1
	}

	for _, row := range grid {
		fmt.Println(row)
	}

	return grid
}

func calculateDots(points []Point) int {
	dots := 0
	grid := getPointGrid(points)

	for x, row := range grid {
		for y, _ := range row {
			if grid[x][y] == 1 {
				dots += 1
			}
		}
	}

	return dots
}
