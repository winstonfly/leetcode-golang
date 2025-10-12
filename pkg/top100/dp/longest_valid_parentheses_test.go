package dp

import "testing"

func TestLongestValidParentheses(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		tests := "()"
		v := longestValidParentheses(tests)
		if v != 2 {
			t.Errorf("longestValidParentheses(%s) = %d, want 2", tests, v)
		}
	})

	t.Run("test2", func(t *testing.T) {
		tests := ")()())"
		v := longestValidParentheses(tests)
		if v != 4 {
			t.Errorf("longestValidParentheses(%s) = %d, want 4", tests, v)
		}
	})

	t.Run("test3", func(t *testing.T) {
		tests := ""
		v := longestValidParentheses(tests)
		if v != 0 {
			t.Errorf("longestValidParentheses(%s) = %d, want 0", tests, v)
		}
	})

	t.Run("test4", func(t *testing.T) {
		tests := "(()"
		v := longestValidParentheses(tests)
		if v != 2 {
			t.Errorf("longestValidParentheses(%s) = %d, want 2", tests, v)
		}
	})

	t.Run("test5", func(t *testing.T) {
		tests := "())"
		v := longestValidParentheses(tests)
		if v != 2 {
			t.Errorf("longestValidParentheses(%s) = %d, want 2", tests, v)
		}
	})
}
