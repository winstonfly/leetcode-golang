package list

import (
	"leetcode-golang/pkg/entity"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindromeList(head *entity.ListNode) bool {
	a := []int{}
	for head != nil {
		a = append(a, head.Val)
		head = head.Next
	}

	i, j := 0, len(a)-1
	for i <= j {
		if a[i] != a[j] {
			return false
		}
		i++
		j--
	}

	return true
}
