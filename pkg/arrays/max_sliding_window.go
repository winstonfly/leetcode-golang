package arrays

import (
	"container/heap"
	"sort"
)

// NO.239
func maxSlidingWindow(nums []int, k int) []int {

	var ans []int

	q := []int{}
	push := func(i int) {
		for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}

	for i := 0; i < k; i++ {
		push(i)
	}

	ans = append(ans, nums[q[0]])

	for i := k; i < len(nums); i++ {
		push(i)
		if q[0] <= i-k {
			q = q[1:]
		}

		ans = append(ans, nums[q[0]])
	}
	return ans
}

var a []int

type hp struct {
	sort.IntSlice
}

func (h hp) Less(i, j int) bool { return a[h.IntSlice[i]] > a[h.IntSlice[j]] } //最大堆

func (h *hp) Push(x any) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *hp) Pop() any {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func maxSlidingWindow1(nums []int, k int) []int {
	var ans []int
	h := &hp{make([]int, k)}
	a = nums

	for i := 0; i < k; i++ {
		h.IntSlice[i] = i
	}
	heap.Init(h)

	ans = append(ans, nums[h.IntSlice[0]])
	for i := k; i < len(nums); i++ {
		heap.Push(h, i)
		for h.IntSlice[0] <= i-k {
			heap.Pop(h)
		}

		ans = append(ans, nums[h.IntSlice[0]])
	}

	return ans
}
