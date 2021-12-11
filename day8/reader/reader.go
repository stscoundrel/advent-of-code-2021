package reader

import (
	"bufio"
	"fmt"
	"os"
)

func ReadSignalInputFromFile(filePath string) (lines []string) {
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Could not read given file: ", filePath)
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
