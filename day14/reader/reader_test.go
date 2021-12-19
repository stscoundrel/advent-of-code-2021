package reader

import (
	"testing"
)

func TestReadsInstructions(t *testing.T) {
	template, rules := ReadInstructions("./test_input.txt")
	expectedTemplate := "NNCB"
	expectedRules := []string{
		"CH -> B",
		"HH -> N",
		"CB -> H",
		"NH -> C",
		"HB -> C",
		"HC -> B",
		"HN -> C",
		"NN -> C",
		"BH -> H",
		"NC -> B",
		"NB -> B",
		"BN -> B",
		"BB -> N",
		"BC -> B",
		"CC -> N",
		"CN -> C",
	}

	for index, rule := range rules {
		if rule != expectedRules[index] {
			t.Error("Did not return expected content. Received", rule, "expected ", expectedRules[index])
		}

	}

	if template != expectedTemplate {
		t.Error("Did not return expected content. Received", template, "expected ", expectedTemplate)
	}

}
