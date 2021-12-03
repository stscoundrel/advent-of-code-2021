package measurements

import (
	"testing"
)

func TestCountsIncreases(t *testing.T) {
	values := []int{
		199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263,
	}

	result := CountValueIncreases(values)

	if result != 7 {
		t.Error("Did not return expected content. Received", result, "expected 7")
	}

}

func TestCountsIncreasesSlidingAverage(t *testing.T) {
	values := []int{
		607,
		618,
		618,
		617,
		647,
		716,
		769,
		792,
	}

	result := CountValueIncreasesSliding(values)
	expected := 5
	if result != expected {
		t.Error("Did not return expected content. Received", result, "expected", expected)
	}

}
