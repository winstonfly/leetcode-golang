package list

import "leetcode-golang/pkg/entity"

func removeNthFromEnd(head *entity.ListNode, n int) *entity.ListNode {
	l := 0
	current := head
	for {
		if current != nil {
			current = current.Next
			l++
		} else {
			break
		}
	}

	//删除第l -n +1个节点
	//创建一个哨兵节点
	dummy := &entity.ListNode{Next: head}
	index := 1
	prev := dummy
	current = dummy.Next
	for index < l-n+1 {
		prev = current
		current = current.Next
		index++
	}

	//删除
	prev.Next = current.Next

	return dummy.Next
}
