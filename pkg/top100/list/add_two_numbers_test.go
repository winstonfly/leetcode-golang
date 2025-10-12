package list

import (
	"slices"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	t.Run("addTwoNumbers", func(t *testing.T) {
		l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
		l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
		result := addTwoNumbers(l1, l2)
		if result.Val != 7 || result.Next.Val != 0 || result.Next.Next.Val != 8 {
			t.Errorf("Expected 7 -> 0 -> 8, but got %d -> %d -> %d", result.Val, result.Next.Val, result.Next.Next.Val)
		}
	})

	t.Run("addTwoNumbers with overflow", func(t *testing.T) {
		l1 := BuildListNode([]int{9, 9, 9, 9, 9, 9, 9})
		l2 := BuildListNode([]int{9, 9, 9, 9})
		result := addTwoNumbers(l1, l2)
		var resultSlice []int
		for result != nil {
			resultSlice = append(resultSlice, result.Val)
			result = result.Next
		}
		expected := []int{8, 9, 9, 9, 0, 0, 0, 1}
		if !slices.Equal(resultSlice, expected) {
			t.Errorf("Expected 8 -> 9 -> 9 -> 9 -> 0 -> 0 -> 0 -> 1, but got %v", resultSlice)
		}

	})
}
