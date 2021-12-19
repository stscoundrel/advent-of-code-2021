package reader

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadInstructions(filePath string) (string, []string) {
	template := ""
	rules := []string{}
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Could not read given file: ", filePath)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) > 0 {
			if strings.Contains(line, " -> ") {
				rules = append(rules, line)
			} else {
				template = line
			}
		}
	}

	file.Close()

	return template, rules
}
