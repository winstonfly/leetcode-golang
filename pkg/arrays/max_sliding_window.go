package arrays

import (
	"container/heap"
	"sort"
)

func maxSlidingWindow1(nums []int, k int) []int {
	var ans []int
	left, right := 1, k
	var maxx, maxxIndex int
	for i := 0; i < k; i++ {
		if nums[i] > maxx {
			maxx = nums[i]
			maxxIndex = i
		}
	}

	ans = append(ans, maxx)
	for left <= len(nums)-k && right < len(nums) {
		if left < maxxIndex {
			maxx = max(maxx, nums[right])

		} else {
			maxx = nums[left]
			for i := left; i <= right; i++ {
				if nums[i] > maxx {
					maxx = max(maxx, nums[i])
					maxxIndex = i
				}
			}
		}
		ans = append(ans, maxx)
		left++
		right++
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

func maxSlidingWindow(nums []int, k int) []int {
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
