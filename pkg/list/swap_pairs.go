package list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs1(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy
	current := dummy.Next
	for current != nil && current.Next != nil {
		//交换过程
		next := current.Next
		//1删除下一个节点
		current.Next = next.Next
		//2将下一个节点插入到当前节点的前面
		next.Next = current
		prev.Next = next

		prev = current
		current = current.Next

	}

	return dummy.Next
}

// 递归解法
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := head.Next
	head.Next = swapPairs(newHead.Next)
	newHead.Next = head

	return newHead
}
