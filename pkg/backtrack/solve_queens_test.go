package backtrack

import "testing"

func TestSolveQueen(t *testing.T) {
	t.Run("solve queens", func(t *testing.T) {
		queens := solveNQueens(1)
		t.Log(queens)
	})
}
