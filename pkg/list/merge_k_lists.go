package list

// NO.23
func mergeKLists1(lists []*ListNode) *ListNode {
	var dummy *ListNode
	for i := 0; i < len(lists); i++ {
		dummy = mergeTwoLists(dummy, lists[i])
	}
	return dummy.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	mid := len(lists) / 2
	left := mergeKLists(lists[:mid])
	right := mergeKLists(lists[mid:])
	return mergeTwoLists(left, right)
}
