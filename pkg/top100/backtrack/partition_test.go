package backtrack

import "testing"

func TestPartition(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		v := partition("aab")
		t.Log(v)
	})
}
