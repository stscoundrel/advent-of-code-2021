package fishcounter

import (
	"testing"
)

func TestCalculatesFishAmount(t *testing.T) {
	initialFishes := []int{
		3, 4, 3, 1, 2,
	}

	fishes1 := NewFishGroup((initialFishes))
	fishes2 := NewFishGroup((initialFishes))

	result1 := GetFishCount(fishes1, 18)
	result2 := GetFishCount(fishes2, 80)

	expected1 := 26
	expected2 := 5934

	if result1 != expected1 {
		t.Error("Did not return expected content. Received", result1, "expected ", expected1)
	}

	if result2 != expected2 {
		t.Error("Did not return expected content. Received", result2, "expected ", expected2)
	}

}
