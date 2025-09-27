package backtrack

import "testing"

func TestExist(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		board := [][]byte{
			{'A', 'B', 'C', 'E'},
			{'S', 'F', 'C', 'S'},
			{'A', 'D', 'E', 'E'},
		}
		word := "ABCCED"
		v := exist(board, word)
		if !v {
			t.Errorf("exist() = %v, want %v", v, true)
		}
	})

	t.Run("test1", func(t *testing.T) {
		board := [][]byte{
			{'A', 'B', 'C', 'E'},
			{'S', 'F', 'C', 'S'},
			{'A', 'D', 'E', 'E'},
		}
		word := "ABCB"
		v := exist(board, word)
		want := false
		if v != want {
			t.Errorf("exist() = %v, want %v", v, want)
		}
	})

	t.Run("test2", func(t *testing.T) {
		board := [][]byte{
			{'a'},
		}
		word := "a"
		v := exist(board, word)
		want := true
		if v != want {
			t.Errorf("exist() = %v, want %v", v, want)
		}
	})

	t.Run("test3", func(t *testing.T) {
		board := [][]byte{
			{'a', 'b'},
			{'c', 'd'},
		}
		word := "abcd"
		v := exist(board, word)
		want := false
		if v != want {
			t.Errorf("exist() = %v, want %v", v, want)
		}
	})

	t.Run("test4", func(t *testing.T) {
		board := [][]byte{
			{'A', 'B', 'C', 'E'},
			{'S', 'F', 'C', 'S'},
			{'A', 'D', 'E', 'E'},
		}
		word := "SEE"
		v := exist(board, word)
		if !v {
			t.Errorf("exist() = %v, want %v", v, true)
		}
	})
}
