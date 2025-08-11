package list

import (
	"fmt"
	"leetcode-golang/pkg/entity"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	head := entity.BuildListNode([]int{1, 1, 2, 1})
	t.Run("is_palindrome", func(t *testing.T) {
		r := isPalindromeList(head)
		fmt.Println(r)
	})
}
