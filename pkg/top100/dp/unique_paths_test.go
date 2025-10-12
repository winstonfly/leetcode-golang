package dp

import "testing"

func TestUniquePaths(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		v := uniquePaths(3, 2)
		t.Log(v)
	})
}
