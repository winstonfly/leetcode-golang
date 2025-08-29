package arrays

import "testing"

func TestFindKthLargest(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		v := findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2)
		if v != 5 {
			t.Errorf("expected 5, got %d", v)
		}

	})
}
