package list

// NO.148
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head

	//找到中间节点
	var prev *ListNode
	for slow != nil && fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next

		if fast == nil {
			break
		}
	}

	mid := prev
	right := mid.Next
	mid.Next = nil
	leftList := sortList(head)
	rightList := sortList(right)

	return mergeTwoLists(leftList, rightList)
}
