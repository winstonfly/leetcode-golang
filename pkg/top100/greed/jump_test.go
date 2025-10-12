package greed

import "testing"

func TestJump(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		v := jump([]int{2, 3, 0, 1, 4})
		t.Log(v)
	})
}
