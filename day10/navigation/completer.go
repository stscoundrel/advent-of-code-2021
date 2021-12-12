package navigation

import "strings"

func GetSequenceCompletion(reading Reading) []string {
	missingEnds := []string{}
	lonelyStarts := strings.Split(reading.debugSign, "")

	for _, start := range lonelyStarts {
		missingEnds = append(missingEnds, openSymbols[start])
	}

	return missingEnds
}
