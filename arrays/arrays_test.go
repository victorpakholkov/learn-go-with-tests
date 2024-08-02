package arrays

import (
	// "reflect"
	"slices"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("array of 9 numbers", func(t *testing.T) {
		numbers := []int{4, 5, 6, 7, 8, 9, 10, 11, 12}
		got := Sum(numbers)
		wanted := 72

		if got != wanted {
			t.Errorf("expected: %d , got: %d, input: %v", wanted, got, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{14, 18}, []int{17, 22})
	wanted := []int{32, 39}

	// !reflect.DeepEqual
	if !slices.Equal(got, wanted) {
		t.Errorf("Error! Wanted: %v, got: %v", wanted, got)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, wanted []int) {
		t.Helper()
		if !slices.Equal(got, wanted) {
			t.Errorf("Error! Wanted: '%v', got: '%v'", wanted, got)
		}
	}
	t.Run("all non-empty", func(t *testing.T) {
		got := SumAllTails([]int{3, 12, 64, 135, 56}, []int{98, 43, 122, 546, 1})
		wanted := []int{267, 712}

		checkSums(t, got, wanted)
	})
	t.Run("includes empty", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{11, 22, 31})
		wanted := []int{0, 53}

		checkSums(t, got, wanted)
	})
}
