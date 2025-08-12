package sort

import "testing"

func TestSortArray(t *testing.T) {

	t.Run("sortArray", func(t *testing.T) {
		nums := []int{5, 2, 3, 1, 4}
		sorted := sortArray(nums)
		t.Log(sorted) // Expected output: [1, 2, 3, 4, 5]
	})
}
