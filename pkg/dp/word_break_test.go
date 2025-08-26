package dp

import "testing"

func TestWordBreak(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		v := wordBreak("leetcode", []string{"leet", "code"})
		if !v {
			t.Errorf("expected true, got false")
		}

	})
}
