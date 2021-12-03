package reader

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ReadIntsFromFile(filePath string) (lines []int) {
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Could not parse min & max length. Check you're using numbers")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, value)
	}

	file.Close()

	return
}
