package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 题号2：两数相加

func buildListNode(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{Val: nums[0]}
	current := head

	for i := 1; i < len(nums); i++ {
		current.Next = &ListNode{Val: nums[i]}
		current = current.Next
	}

	return head
}

func main() {
	addTwoNumbers(buildListNode([]int{9, 9, 9, 9, 9, 9, 9}), buildListNode([]int{9, 9, 9, 9}))
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return nil
}
