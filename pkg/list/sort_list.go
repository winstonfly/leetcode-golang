package list

func sortList(head *ListNode) *ListNode {
	var tmp []int
	for head != nil {
		tmp = append(tmp, head.Val)
		head = head.Next
	}

	result := sortArray(tmp)
	dummy := &ListNode{}
	pre := dummy
	for _, v := range result {
		current := &ListNode{Val: v}
		pre.Next = current
		pre = current
	}

	return dummy.Next
}

func sortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	left := sortArray(nums[0 : len(nums)/2])
	right := sortArray(nums[len(nums)/2:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	if i < len(left) {
		result = append(result, left[i:]...)
	}

	if j < len(right) {
		result = append(result, right[j:]...)
	}

	return result
}
