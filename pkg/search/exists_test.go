package search

import "testing"

func TestExists(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		board := [][]byte{
			{'A', 'B', 'C', 'E'},
			{'S', 'F', 'C', 'S'},
			{'A', 'D', 'E', 'E'},
		}
		word := "ABCCED"
		result := exist(board, word)
		if result != true {
			t.Errorf("exists(%q, %q) = %t; want %t", board, word, result, true)
		}
	})

	t.Run("test2", func(t *testing.T) {
		board := [][]byte{
			{'A', 'B', 'C', 'E'},
			{'S', 'F', 'C', 'S'},
			{'A', 'D', 'E', 'E'},
		}
		word := "SEE"
		result := exist(board, word)
		if result != true {
			t.Errorf("exists(%q, %q) = %t; want %t", board, word, result, true)
		}
	})

	t.Run("test3", func(t *testing.T) {
		board := [][]byte{
			{'A', 'B', 'C', 'E'},
			{'S', 'F', 'C', 'S'},
			{'A', 'D', 'E', 'E'},
		}
		word := "ABCB"
		result := exist(board, word)
		if result != true {
			t.Errorf("exists(%q, %q) = %t; want %t", board, word, result, true)
		}
	})

}
