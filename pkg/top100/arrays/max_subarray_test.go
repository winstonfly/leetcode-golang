package arrays

import "testing"

func TestMaxSubArray(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		tests := []int{2, 1, -3, 4, -1, 2, 1, -5, 4}
		v := maxSubArray(tests)
		if v != 6 {
			t.Errorf("expected 6, got %d", v)
		}
	})
}
