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

	fmt.Println(firstLine)

	file.Close()

	return
}
