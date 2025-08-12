package list

func mergeKLists(lists []*ListNode) *ListNode {
	var dummy *ListNode
	for i := 0; i < len(lists); i++ {
		dummy = mergeTwoLists(dummy, lists[i])
	}
	return dummy.Next
}

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	ll, lr := l1, l2
	dummy := &ListNode{}
	current := dummy
	for ll != nil && lr != nil {
		if ll.Val < lr.Val {
			current.Next = ll
			ll = ll.Next
		} else {
			current.Next = lr
			lr = lr.Next
		}
		current = current.Next
	}

	if ll != nil {
		current.Next = ll
	}

	if lr != nil {
		current.Next = lr
	}

	return dummy.Next
}
