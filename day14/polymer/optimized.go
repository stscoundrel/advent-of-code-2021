package polymer

import "fmt"

func getRule(pair string, rules []Rule) Rule {
	result := Rule{}

	for _, rule := range rules {
		if rule.pair == pair {
			result = rule
		}
	}

	return result
}

func getInitialMap(template string, rules []Rule) map[string]int {
	frequencyMap := map[string]int{}

	for i := 0; i < len(template)-1; i += 1 {
		pair := template[i : i+2]
		rule := getRule(pair, rules)
		firstPair := string(template[i]) + rule.value
		secondPair := rule.value + string(template[i+1])

		frequencyMap[firstPair]++
		frequencyMap[secondPair]++

	}

	return frequencyMap
}

func GetMostAndLeastCommon(template string, rules []Rule, iterations int) (int, int) {
	frequencyMap := getInitialMap(template, rules)
	fmt.Println(frequencyMap)
	for step := 1; step < iterations; step += 1 {
		newMap := map[string]int{}

		for pair, count := range frequencyMap {
			rule := getRule(pair, rules)

			firstPair := string(pair[0]) + rule.value
			secondPair := rule.value + string(pair[1])

			newMap[firstPair] += count
			newMap[secondPair] += count
		}

		frequencyMap = newMap
		fmt.Println(frequencyMap)
	}

	// Throw in the missing last letter, for decencys sake.
	frequencyMap[string(template[len(template)-1])]++

	return getMostAndLeastCommonInPairMap(frequencyMap)
}
