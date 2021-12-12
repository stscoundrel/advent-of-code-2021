package navigation

import (
	"strings"
	"testing"
)

func TestRemovesLastMatchOf(t *testing.T) {
	lines := []string{
		"[",
		"[",
		"(",
		"{",
		"<",
	}

	result1 := removeLastMatchOf(">", lines)
	result2 := removeLastMatchOf("}", lines)
	result3 := removeLastMatchOf(")", lines)

	expected1 := "[[({"
	expected2 := "[[(<"
	expected3 := "[[{<"

	if strings.Join(result1, "") != expected1 {
		t.Error("Result did not match expected content. Got", strings.Join(result1, ""), "expected ", expected1)
	}

	if strings.Join(result2, "") != expected2 {
		t.Error("Result did not match expected content. Got", strings.Join(result2, ""), "expected ", expected2)
	}

	if strings.Join(result3, "") != expected3 {
		t.Error("Result did not match expected content. Got", strings.Join(result3, ""), "expected ", expected3)
	}
}
