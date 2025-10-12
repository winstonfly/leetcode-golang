package dp

import "testing"

func TestLongestCommanSubsequence(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		v := longestCommonSubsequence("abcde", "ace")
		want := 3
		if v != want {
			t.Errorf("longestCommonSubsequence() = %d, want %d", v, want)
		}
	})
}
