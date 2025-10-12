package arrays

import "testing"

func TestTrap(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		v := trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
		t.Log(v)
	})
}
