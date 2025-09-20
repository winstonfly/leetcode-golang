package list

import "testing"

func TestReverseList(t *testing.T) {

	t.Run("test1", func(t *testing.T) {
		head := &ListNode{Val: 1}
		head.Next = &ListNode{Val: 2}
		head.Next.Next = &ListNode{Val: 3}
		head.Next.Next.Next = &ListNode{Val: 4}
		head.Next.Next.Next.Next = &ListNode{Val: 5}

		newHead := reverseList(head)
		for current := newHead; current != nil; current = current.Next {
			t.Log(current.Val)
		}
	})
}
