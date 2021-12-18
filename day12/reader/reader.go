package reader

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadPathsFromFile(filePath string) map[string][]string {
	paths := map[string][]string{}
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Could not read given file: ", filePath)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "-")

		_, hasFirstPart := paths[parts[0]]
		if !hasFirstPart {
			paths[parts[0]] = []string{}
		}

		paths[parts[0]] = append(paths[parts[0]], parts[1])

		_, hasSecondPart := paths[parts[1]]
		if !hasSecondPart {
			paths[parts[1]] = []string{}
		}
		paths[parts[1]] = append(paths[parts[1]], parts[0])
	}

	file.Close()

	return paths
}
