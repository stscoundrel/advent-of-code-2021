package reader

import (
	"testing"
)

func TestReadsTransmissionFromFile(t *testing.T) {
	result1 := ReadGTransmissionFromFile("./test_input.txt")
	expected1 := "8A004A801A8002F478"

	result2 := ReadGTransmissionFromFile("./test_input2.txt")
	expected2 := "620080001611562C8802118E34"

	result3 := ReadGTransmissionFromFile("./test_input3.txt")
	expected3 := "C0015000016115A2E0802F182340"

	result4 := ReadGTransmissionFromFile("./test_input4.txt")
	expected4 := "A0016C880162017C3686B18A3D4780"

	if result1 != expected1 {
		t.Error("Did not return expected content. Received", result1, "expected ", expected1)
	}

	if result2 != expected2 {
		t.Error("Did not return expected content. Received", result2, "expected ", expected2)
	}

	if result3 != expected3 {
		t.Error("Did not return expected content. Received", result3, "expected ", expected3)
	}

	if result4 != expected4 {
		t.Error("Did not return expected content. Received", result4, "expected ", expected4)
	}

}
