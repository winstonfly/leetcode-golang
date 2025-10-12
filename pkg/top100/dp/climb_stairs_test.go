package dp

import "testing"

func TestClimbStairs(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		v := climbStairs(2)
		t.Log(v)
	})
}
