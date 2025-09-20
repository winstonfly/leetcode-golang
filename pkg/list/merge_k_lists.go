package list

func mergeKLists(lists []*ListNode) *ListNode {
	var dummy *ListNode
	for i := 0; i < len(lists); i++ {
		dummy = mergeTwoLists(dummy, lists[i])
	}
	return dummy.Next
}
