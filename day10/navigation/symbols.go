package navigation

var closeSymbols map[string]string = map[string]string{
	")": "(",
	"]": "[",
	"}": "{",
	">": "<",
}

func isStartSymbol(symbol string) bool {
	if symbol == "(" || symbol == "[" || symbol == "{" || symbol == "<" {
		return true
	}

	return false
}

func matchesStartSymbol(close string, start string) bool {
	return closeSymbols[close] == start
}
