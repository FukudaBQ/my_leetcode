package leetcode347

import "container/heap"

type item struct {
	val  int
	freq int
}
type minHeap []item

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	return h[i].freq < h[j].freq
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(x any) {
	*h = append(*h, x.(item))
}

func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func topKFrequentHeap(nums []int, k int) []int {
	freq := make(map[int]int)
	for _, n := range nums {
		freq[n]++
	}
	h := &minHeap{}
	heap.Init(h)
	for v, f := range freq {
		if h.Len() < k {
			heap.Push(h, item{v, f})
		} else if f > (*h)[0].freq {
			(*h)[0] = item{v, f}
			heap.Fix(h, 0)
		}
	}
	res := make([]int, k)
	for i := k - 1; i >= 0; i-- {
		res[i] = heap.Pop(h).(item).val
	}
	return res
}
