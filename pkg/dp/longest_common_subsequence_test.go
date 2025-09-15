package dp

import "testing"

func TestLongestCommanSubsequence(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		v := longestCommonSubsequence("abcde", "ace")
		t.Log(v)
		if v != 3 {
			t.Errorf("expected 3, got %d", v)
		}
	})
}
