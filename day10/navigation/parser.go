package navigation

import "strings"

func normalizeInput(input string) string {
	normalizedInput := input
	foundReplaces := true

	knownSingles := [4]string{
		"()",
		"[]",
		"{}",
		"<>",
	}

	for foundReplaces {
		previousInput := normalizedInput
		for _, knownSingle := range knownSingles {
			normalizedInput = strings.Replace(normalizedInput, knownSingle, "", -1)
		}

		if previousInput == normalizedInput {
			foundReplaces = false
		}
	}

	return normalizedInput
}
