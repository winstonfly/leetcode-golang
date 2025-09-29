package stacks

import "testing"

func TestDecodeString(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		v := decodeString("3[a]2[bc]")
		want := "aaabcbc"
		if v != want {
			t.Errorf("decodeString failed, want:%v, got:%v", want, v)
		}
	})

	t.Run("test2", func(t *testing.T) {
		v := decodeString("100[leetcode]")
		want := "accaccacc"
		if v != want {
			t.Errorf("decodeString failed, want:%v, got:%v", want, v)
		}
	})
}
