package main

import "fmt"

// "awekpek"

func main() {
	result := lengthOfLongestSubstring("awekpek")
	fmt.Println(result)
}

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int, 1)

	left, right, maxLength := 0, 0, 0

	for right < len(s) {
		if index, found := m[s[right]]; found && index >= left {
			left = index + 1
		}

		m[s[right]] = right
		maxLength = max(maxLength, right-left+1)

		right++
	}

	return maxLength
}
