package reader

import (
	"bufio"
	"fmt"
	"os"
)

func ReadStringsFromFile(filePath string) (lines []string) {
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Could read the file with path ", filePath)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	file.Close()

	return
}
