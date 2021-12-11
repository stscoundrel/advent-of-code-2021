package heights

import (
	"testing"
)

func TestGetsRisklevel(t *testing.T) {
	lowpoints := []int{
		1, 0, 5, 5,
	}

	result := CalculateRiskLevel(lowpoints)
	expected := 15

	if result != expected {
		t.Error("Did not return expected content. Received", result, "expected ", expected)
	}

}
