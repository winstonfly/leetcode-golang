package dp

import "testing"

func TestNumSqure(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		n := 12
		v := numSquares(n)
		if v != 3 {
			t.Errorf("expected 3, got %d", v)
		}
	})
}
