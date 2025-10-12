package arrays

import "testing"

func TestDecodeString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"3[a]2[bc]", "aaabcbc"},
		{"3[a2[c]]", "accaccacc"},
		{"2[abc]3[cd]ef", "abcabccdcdcdef"},
		{"abc3[cd]xyz", "abccdcdcdxyz"},
		{"2[3[a]b]", "aaabaaab"},
	}

	for _, test := range tests {
		result := decodeString(test.input)
		if result != test.expected {
			t.Errorf("decodeString(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}
