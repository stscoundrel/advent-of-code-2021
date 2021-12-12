package navigation

import (
	"testing"
)

func TestGetsReadingStatus(t *testing.T) {
	reading1 := NewReading("{([(<{}[<>[]}>{[]{[(<()>")
	reading2 := NewReading("[[<[([]))<([[{}[[()]]]")

	if reading1.status != CORRUPTED {
		t.Error("Result did not match expected output. Got", reading1.status, "expected ", CORRUPTED)
	}

	if reading1.debugSign != "}" {
		t.Error("Result did not match expected output. Got", reading1.debugSign, "expected }")
	}

	if reading2.status != CORRUPTED {
		t.Error("Result did not match expected output. Got", reading2.status, "expected ", CORRUPTED)
	}

	if reading2.debugSign != ")" {
		t.Error("Result did not match expected output. Got", reading2.debugSign, "expected )")
	}
}

func TestGetsReadingsWithCorruptedStatus(t *testing.T) {
	lines := []string{
		"[({(<(())[]>[[{[]{<()<>>",
		"[(()[<>])]({[<{<<[]>>(",
		"{([(<{}[<>[]}>{[]{[(<()>",
		"(((({<>}<{<{<>}{[]{[]{}",
		"[[<[([]))<([[{}[[()]]]",
		"[{[{({}]{}}([{[{{{}}([]",
		"{<[[]]>}<{[{[{[]{()[[[]",
		"[<(<(<(<{}))><([]([]()",
		"<{([([[(<>()){}]>(<<{{",
		"<{([{{}}[<[[[<>{}]]]>[]]",
	}

	readings := NewReadings(lines)
	corruptedReadings := GetCorruptedReadings(readings)

	expected := []string{
		"{([(<{}[<>[]}>{[]{[(<()>",
		"[[<[([]))<([[{}[[()]]]",
		"[{[{({}]{}}([{[{{{}}([]",
		"[<(<(<(<{}))><([]([]()",
		"<{([([[(<>()){}]>(<<{{ ",
	}

	if len(corruptedReadings) != len(expected) {
		t.Error("Result did not match expected length. Got", len(corruptedReadings), "expected ", len(expected))
	}
}

func TestGetsReadingsWithIncompleteStatus(t *testing.T) {
	lines := []string{
		"[({(<(())[]>[[{[]{<()<>>",
		"[(()[<>])]({[<{<<[]>>(",
		"{([(<{}[<>[]}>{[]{[(<()>",
		"(((({<>}<{<{<>}{[]{[]{}",
		"[[<[([]))<([[{}[[()]]]",
		"[{[{({}]{}}([{[{{{}}([]",
		"{<[[]]>}<{[{[{[]{()[[[]",
		"[<(<(<(<{}))><([]([]()",
		"<{([([[(<>()){}]>(<<{{",
		"<{([{{}}[<[[[<>{}]]]>[]]",
	}

	readings := NewReadings(lines)
	incompleteReadings := GetIncompleteReadings(readings)

	expected := []string{
		"[({(<(())[]>[[{[]{<()<>>", // Complete by adding }}]])})].
		"[(()[<>])]({[<{<<[]>>(",   // Complete by adding )}>]}).
		"(((({<>}<{<{<>}{[]{[]{}",  // Complete by adding }}>}>)))).
		"{<[[]]>}<{[{[{[]{()[[[]",  // Complete by adding ]]}}]}]}>.
		"<{([{{}}[<[[[<>{}]]]>[]]", // Complete by adding ])}>.
	}

	if len(incompleteReadings) != len(expected) {
		t.Error("Result did not match expected length. Got", len(incompleteReadings), "expected ", len(expected))
	}
}
