package backtrack

import "testing"

func TestSolveQueen(t *testing.T) {
	t.Run("solve queens", func(t *testing.T) {
		queens := solveNQueens(4)
		t.Log(queens)
	})

	t.Run("solve queens2", func(t *testing.T) {
		queens := solveNQueens(2)
		t.Log(queens)
	})

	t.Run("solve queens3", func(t *testing.T) {
		queens := solveNQueens(5)
		for _, v := range queens {
			t.Log(v)
		}
	})

}
