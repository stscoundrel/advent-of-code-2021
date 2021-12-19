package reader

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadFoldInstructions(filePath string) ([][]int, []string) {
	points := [][]int{}
	folds := []string{}
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Could not read given file: ", filePath)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) > 0 {
			if strings.Contains(line, "fold") {
				folds = append(folds, line)
			} else {
				pointParts := strings.Split(line, ",")
				x, _ := strconv.Atoi(pointParts[0])
				y, _ := strconv.Atoi(pointParts[1])

				point := []int{x, y}

				points = append(points, point)
			}
		}
	}

	file.Close()

	return points, folds
}
