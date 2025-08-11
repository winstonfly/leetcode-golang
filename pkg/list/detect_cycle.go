package list

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	if head.Next == head {
		return head
	}

	slow := head
	fast := head.Next
	index := 0
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return nil
		}

		if slow == slow.Next {
			return slow
		}
		slow = slow.Next
		fast = fast.Next.Next
		index++
	}

	i := 1
	dummy := &ListNode{Next: head}
	current := dummy
	for current.Next != nil && i < index {
		current = current.Next
		i++
	}

	return current
}
