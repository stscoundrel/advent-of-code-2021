package reader

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadHeighMapFromFile(filePath string) [][]int {
	heights := [][]int{}
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Could read the file with path ", filePath)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		rowNumbers := []int{}
		values := strings.Split(line, "")

		for _, value := range values {
			if len(value) > 0 {
				number, _ := strconv.Atoi(value)
				rowNumbers = append(rowNumbers, number)
			}
		}

		heights = append(heights, rowNumbers)
	}

	file.Close()

	return heights
}
