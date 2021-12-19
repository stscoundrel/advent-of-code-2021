package polymer

import "strings"

type Rule struct {
	pair   string
	value  string
	first  string
	second string
}

func newRule(rule string) Rule {
	parts := strings.Split(rule, " -> ")
	pairParts := strings.Split(parts[0], "")

	return Rule{pair: parts[0], value: parts[1], first: pairParts[0], second: pairParts[1]}
}

func ParseRules(rawRules []string) []Rule {
	rules := []Rule{}

	for _, rawRule := range rawRules {
		rules = append(rules, newRule(rawRule))
	}

	return rules
}
