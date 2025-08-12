package list

func sortList(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	//slow是中间节点
	//left := head, right := slow.Next

	return nil
}

func mergeSqeuence(left, right *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for {
		if left == nil {
			current.Next = right
			return dummy.Next
		}

		if right == nil {
			current.Next = left
			return dummy.Next
		}

		if left.Val < right.Val {
			current.Next = left
			left = left.Next
		} else {
			current.Next = right
			right = right.Next
		}

		current = current.Next
	}

	return dummy.Next
}
