package median

import (
	"container/heap"
)

// A MinIntHeap is a min-heap of ints.
type MinIntHeap []int

func (h MinIntHeap) Len() int           { return len(h) }
func (h MinIntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MinIntHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h *MinIntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))

	heap.Fix(h, 0)
}

func (h *MinIntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	heap.Fix(h, 0)
	return x
}
