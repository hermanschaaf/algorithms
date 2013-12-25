package median

import (
	"container/heap"
)

// A MaxIntHeap is a max-heap of ints.
type MaxIntHeap []int

func (h MaxIntHeap) Len() int           { return len(h) }
func (h MaxIntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MaxIntHeap) Less(i, j int) bool { return h[j] < h[i] }

func (h *MaxIntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
	heap.Fix(h, 0)
}

func (h *MaxIntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	heap.Fix(h, 0)
	return x
}
