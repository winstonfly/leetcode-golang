package arrays

import "container/heap"

type Pair struct {
	key   int
	value int
}

type pairHeap struct {
	pairs []Pair
}

func (h pairHeap) Len() int {
	return len(h.pairs)
}

// 大顶堆
func (h pairHeap) Less(i, j int) bool {
	return h.pairs[i].value > h.pairs[j].value
}

func (h pairHeap) Swap(i, j int) {
	h.pairs[i], h.pairs[j] = h.pairs[j], h.pairs[i]
}

func (h *pairHeap) Push(x interface{}) {
	h.pairs = append(h.pairs, x.(Pair))
}

func (h *pairHeap) Pop() interface{} {
	old := h.pairs
	n := len(old)
	x := old[n-1]
	h.pairs = old[0 : n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]Pair, 1)
	for i := 0; i < len(nums); i++ {
		if p, ok := m[nums[i]]; ok {
			p.value++
			m[nums[i]] = p
		} else {
			m[nums[i]] = Pair{nums[i], 1}
		}
	}

	ph := &pairHeap{
		pairs: make([]Pair, 0, len(m)),
	}
	for _, v := range m {
		ph.pairs = append(ph.pairs, v)
	}

	heap.Init(ph)

	var ans []int
	for i := 0; i < k; i++ {
		ans = append(ans, ph.pairs[0].key)
		heap.Pop(ph)
	}

	return ans
}
