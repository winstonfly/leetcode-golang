package graph

import "testing"

func TestOrangesRotting(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		tests := [][]int{
			{2, 1, 1},
			{1, 1, 0},
			{0, 1, 1},
		}

		v := orangesRotting(tests)
		if v != 4 {
			t.Errorf("expected 4, got %d", v)
		}
	})
}
