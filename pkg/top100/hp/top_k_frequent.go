package hp

import "container/heap"

type keyPair struct {
	key   int //值
	value int //数量
}

type pairHeap struct {
	pairs []keyPair
}

func (p *pairHeap) Len() int {
	return len(p.pairs)
}

func (p *pairHeap) Less(i, j int) bool {
	return p.pairs[i].value > p.pairs[j].value
}

func (p *pairHeap) Swap(i, j int) {
	p.pairs[i], p.pairs[j] = p.pairs[j], p.pairs[i]
}

func (p *pairHeap) Push(x any) {
	p.pairs = append(p.pairs, x.(keyPair))
}

func (p *pairHeap) Pop() any {
	v := p.pairs[len(p.pairs)-1]
	p.pairs = p.pairs[:len(p.pairs)-1]
	return v
}

// NO.347
func topKFrequent(nums []int, k int) []int {
	var ans []int
	freqMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		freqMap[nums[i]]++
	}
	var pairs []keyPair
	for i, v := range freqMap {
		pair := keyPair{key: i, value: v}
		pairs = append(pairs, pair)
	}

	h := &pairHeap{pairs: pairs}
	heap.Init(h)

	for i := 0; i < k; i++ {
		ans = append(ans, h.pairs[0].key)
		heap.Pop(h)
	}

	return ans
}
