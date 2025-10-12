package sort

import "testing"

func TestSortColors(t *testing.T) {
	t.Run("sortColors", func(t *testing.T) {
		nums := []int{2, 0, 2, 1, 1, 0}
		sortColors(nums)
		//expected := []int{0, 0, 1, 1, 2, 2}
		t.Log(nums)
	})
}
