package reader

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadIntsFromFile(filePath string) (values []int) {
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Could not read given file: ", filePath)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		numbersInLine := strings.Split(scanner.Text(), ",")

		for _, number := range numbersInLine {
			value, _ := strconv.Atoi(number)
			values = append(values, value)
		}
	}

	file.Close()

	return
}
