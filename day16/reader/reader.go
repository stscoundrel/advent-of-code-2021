package reader

import (
	"bufio"
	"fmt"
	"os"
)

func ReadGTransmissionFromFile(filePath string) string {
	result := ""
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Could not read given file: ", filePath)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		result = scanner.Text()
	}

	file.Close()

	return result
}
