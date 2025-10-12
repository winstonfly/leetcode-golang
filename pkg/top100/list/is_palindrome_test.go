package list

import (
	"fmt"
	"leetcode-golang/pkg/top100/entity"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	t.Run("is_palindrome", func(t *testing.T) {
		head := entity.BuildListNode([]int{1, 1, 2, 1})
		r := isPalindromeList(head)
		fmt.Println(r)
	})

	t.Run("is_palindrome", func(t *testing.T) {
		head := entity.BuildListNode([]int{1, 2, 2, 1})
		r := isPalindromeList(head)
		fmt.Println(r)
	})
}
