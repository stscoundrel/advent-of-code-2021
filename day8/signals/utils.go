package signals

import (
	"sort"
	"strings"
)

func normalizeAndReverse(mapping map[int]string) map[string]int {
	newMap := map[string]int{}

	for key, value := range mapping {
		newMap[sortDigits(value)] = key
	}

	return newMap
}

func sortDigits(digits string) string {
	digitParts := strings.Split(digits, "")
	sort.Strings(digitParts)
	return strings.Join(digitParts, "")
}
