package graph

import "testing"

func TestNumIslands(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		tests := [][]byte{
			{'1', '1', '1', '1', '0'},
			{'1', '1', '0', '1', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '0', '0', '0'},
		}

		v := numIslands(tests)
		if v != 1 {
			t.Errorf("expected 1, got %d", v)
		}
	})

	t.Run("test2", func(t *testing.T) {
		tests := [][]byte{
			{'1'},
			{'1'},
		}

		v := numIslands(tests)
		if v != 1 {
			t.Errorf("expected 1, got %d", v)
		}
	})
}
