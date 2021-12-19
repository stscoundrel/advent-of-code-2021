package polymer

func applyRules(template string, rules []Rule) string {
	newTemplate := ""

	for i := 0; i < len(template); i += 1 {
		if i+1 != len(template) {
			pair := template[i : i+2]
			for _, rule := range rules {
				if rule.pair == pair {
					newTemplate = newTemplate + string(pair[0]) + rule.value
				}
			}
		} else {
			newTemplate = newTemplate + string(template[i])
		}
	}

	return newTemplate
}

func buildPolymer(template string, rules []Rule, iterations int) string {
	buildTemplate := template

	for iteration := 0; iteration < iterations; iteration += 1 {
		buildTemplate = applyRules(buildTemplate, rules)
	}

	return buildTemplate
}

func buildFrequencyMap(template string) map[string]int {
	frequency := map[string]int{}

	for _, letter := range template {
		frequency[string(letter)] += 1
	}

	return frequency
}

func GetMostAndLeastCommonLinear(template string, rules []Rule, iterations int) (int, int) {
	buildTemplate := buildPolymer(template, rules, iterations)
	frequencyMap := buildFrequencyMap(buildTemplate)

	return getMostAndLeastCommonInMap(frequencyMap)
}
