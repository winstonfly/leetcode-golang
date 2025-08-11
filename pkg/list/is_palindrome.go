package list

import (
	"leetcode-golang/pkg/entity"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindromeList(head *entity.ListNode) bool {
	original := head
	list1 := reverseList(head)
	for {
		if list1 == nil || original == nil {
			break
		}
		if list1.Val != original.Val {
			return false
		}

		list1 = list1.Next
		original = original.Next
	}

	return true
}

func reverseList(head *entity.ListNode) *entity.ListNode {
	//创建哨兵节点
	dummy := &entity.ListNode{Next: head}
	var prev *entity.ListNode
	for dummy.Next != nil {
		//获取一个节点
		current := dummy.Next
		dummy.Next = current.Next //删除这个节点
		current.Next = prev
		prev = current //将当前节点放到前面
	}

	return prev
}
