package transmissions

import "testing"

func TestSumsVersionNumbers(t *testing.T) {
	result1 := SumVersionNumbers("8A004A801A8002F478")
	expected1 := 16

	result2 := SumVersionNumbers("620080001611562C8802118E34")
	expected2 := 12

	result3 := SumVersionNumbers("C0015000016115A2E0802F182340")
	expected3 := 23

	result4 := SumVersionNumbers("A0016C880162017C3686B18A3D4780")
	expected4 := 31

	if result1 != expected1 {
		t.Error("Result 1 did not return expected content. Received", result1, "expected ", expected1)
	}

	if result2 != expected2 {
		t.Error("Result 2 did not return expected content. Received", result2, "expected ", expected2)
	}

	if result3 != expected3 {
		t.Error("Result 3 did not return expected content. Received", result3, "expected ", expected3)
	}

	if result4 != expected4 {
		t.Error("Result 4 did not return expected content. Received", result4, "expected ", expected4)
	}
}
