package reader

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadGridFromFile(filePath string) [][]int {
	grid := [][]int{}
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Could not read given file: ", filePath)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	rowCount := 0

	for scanner.Scan() {
		row := []int{}
		numbers := strings.Split(scanner.Text(), "")

		for _, number := range numbers {
			integer, _ := strconv.Atoi(number)
			row = append(row, integer)
		}

		grid = append(grid, row)
		rowCount++
	}

	file.Close()

	return grid
}
