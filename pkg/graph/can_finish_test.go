package graph

import "testing"

func TestCanFinish(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		tests := [][]int{
			{1, 0},
		}

		v := canFinish(2, tests)
		if v != true {
			t.Errorf("expected true, got %t", v)
		}
	})

	t.Run("test", func(t *testing.T) {
		tests := [][]int{
			{1, 0},
			{0, 1},
		}

		v := canFinish(2, tests)
		if v != false {
			t.Errorf("expected false, got %t", v)
		}
	})
}
