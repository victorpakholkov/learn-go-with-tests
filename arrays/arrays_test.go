package arrays

import "testing"

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
