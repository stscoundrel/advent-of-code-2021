package navigation

import (
	"testing"
)

// func TestGetsReadingStatus(t *testing.T) {
// 	reading1 := NewReading("{([(<{}[<>[]}>{[]{[(<()>")
// 	reading2 := NewReading("[[<[([]))<([[{}[[()]]]")

// 	if reading1.status != CORRUPTED {
// 		t.Error("Result did not match expected output. Got", reading1.status, "expected ", CORRUPTED)
// 	}

// 	if reading1.errorSign != "}" {
// 		t.Error("Result did not match expected output. Got", reading1.errorSign, "expected }")
// 	}

// 	if reading2.status != CORRUPTED {
// 		t.Error("Result did not match expected output. Got", reading2.status, "expected ", CORRUPTED)
// 	}

// 	if reading2.errorSign != ")" {
// 		t.Error("Result did not match expected output. Got", reading2.errorSign, "expected )")
// 	}
// }

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
