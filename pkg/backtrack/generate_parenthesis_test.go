package backtrack

import "testing"

func TestGenerateParenthesis(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		v := generateParenthesis(3)
		t.Log(v)
	})
}
