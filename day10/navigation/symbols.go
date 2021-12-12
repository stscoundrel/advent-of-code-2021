package navigation

var closeSymbols map[string]string = map[string]string{
	")": "(",
	"]": "[",
	"}": "{",
	">": "<",
}

var openSymbols map[string]string = map[string]string{
	"(": ")",
	"[": "]",
	"{": "}",
	"<": ">",
}

func isStartSymbol(symbol string) bool {
	if symbol == "(" || symbol == "[" || symbol == "{" || symbol == "<" {
		return true
	}

	return false
}

func isCloseSymbol(symbol string) bool {
	if symbol == ")" || symbol == "]" || symbol == "}" || symbol == ">" {
		return true
	}

	return false
}

func matchesStartSymbol(close string, start string) bool {
	return closeSymbols[close] == start
}

func reverseOpens(opens []string) {
	for i, j := 0, len(opens)-1; i < j; i, j = i+1, j-1 {
		opens[i], opens[j] = opens[j], opens[i]
	}
}

func removeLastMatchOf(close string, opens []string) []string {
	filteredOpens := []string{}
	open := closeSymbols[close]
	hasFound := false

	for i := len(opens) - 1; i >= 0; i -= 1 {
		if !hasFound && opens[i] == open {
			hasFound = true
			continue
		}

		filteredOpens = append(filteredOpens, opens[i])

	}

	reverseOpens(filteredOpens)

	return filteredOpens
}
