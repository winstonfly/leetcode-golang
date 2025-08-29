package arrays

import "testing"

func TestTopKFrquent(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		v := topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)
		if v[0] != 1 || v[1] != 2 {
			t.Errorf("expected [1, 2], got %v", v)
		}

	})
}
