package reader

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadBingoNumbersFromFile(filePath string) (numbers []int) {
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Could read the file with path ", filePath)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var firstLine string

	for scanner.Scan() {
		firstLine = scanner.Text()
		break
	}

	for _, letter := range strings.Split(firstLine, ",") {
		numberValue, _ := strconv.Atoi(string(letter))
		numbers = append(numbers, numberValue)
	}

	file.Close()

	return
}

func ReadBingoBoardsFromFile(filePath string) [][][]int {
	boards := [][][]int{}
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Could read the file with path ", filePath)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	count := 0
	boardsCount := 0
	for scanner.Scan() {
		if count < 2 {
			count++
			continue
		}

		line := scanner.Text()

		if len(line) > 0 {
			row := []int{}
			values := strings.Split(line, " ")

			for _, value := range values {
				if len(value) > 0 {
					number, _ := strconv.Atoi(value)
					row = append(row, number)

				}
			}

			if len(boards) < boardsCount+1 {
				initialBoard := [][]int{row}
				boards = append(boards, initialBoard)
			} else {
				boards[boardsCount] = append(boards[boardsCount], row)
			}

		} else {
			boardsCount++
		}

		count++

	}

	file.Close()

	return boards
}
