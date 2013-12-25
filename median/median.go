package median

import (
	"container/heap"
	"fmt"
)

// Median calculates and stores the running median of a stream
// of numbers
type StreamingMedian struct {
	rightHeap   *MinIntHeap
	leftHeap    *MaxIntHeap
	Median      float64
	Size        int
	initialized bool
}

// Add adds a number to the median stream
// The streaming median is computed using two heaps. All the numbers less than or equal
// to the current median are in the left heap, which is arranged so that the maximum number
// is at the root of the heap. All the numbers greater than or equal to the current median
// are in the right heap, which is arranged so that the minimum number is at the root of
// the heap. Note that numbers equal to the current median can be in either heap.
// The count of numbers in the two heaps never differs by more than 1.
func (m *StreamingMedian) Add(i int) (median float64) {
	// When the process begins the two heaps are initially empty.
	if !m.initialized {
		m.rightHeap = &MinIntHeap{}
		m.leftHeap = &MaxIntHeap{}
		heap.Init(m.rightHeap)
		heap.Init(m.leftHeap)
		m.Size = 0
		m.initialized = true
	}
	if m.Size == 0 {
		// The first number in the input sequence is added to one of the heaps,
		// it doesn’t matter which, and returned as the first streaming median.
		heap.Push(m.leftHeap, i)
		m.Median = float64(i)
		fmt.Println("size 0")
	} else if m.Size == 1 {
		// The second number in the input sequence is then added to the other heap,
		// if the root of the right heap is less than the root of the left heap the
		// two heaps are swapped, and the average of the two numbers is returned as
		// the second streaming median.
		if (*m.leftHeap)[0] > i {
			n := m.leftHeap.Pop()
			heap.Push(m.leftHeap, i)
			heap.Push(m.rightHeap, n)
		} else {
			heap.Push(m.rightHeap, i)
		}
		m.Median = float64((*m.rightHeap)[0]+i) / float64(2.0)
		fmt.Println("size 1")
	} else {
		fmt.Println("main")
		// Now the main algorithm begins. Each subsequent number in the input sequence
		// is compared to the current median, and added to the left heap if it is less
		// than the current median or to the right heap if it is greater than the current
		// median; if the input number is equal to the current median, it is added to
		// whichever heap has the smaller count, or to either heap arbitrarily if they
		// have the same count.
		fi := float64(i)
		if fi < m.Median {
			m.leftHeap.Push(i)
		} else if fi == m.Median {
			if len((*m.rightHeap)) <= len((*m.leftHeap)) {
				m.rightHeap.Push(i)
			} else {
				m.leftHeap.Push(i)
			}
		} else {
			m.rightHeap.Push(i)
		}

		// If that causes the counts of the two heaps to differ by
		// more than 1, the root of the larger heap is removed and inserted in the smaller
		// heap.
		if len((*m.rightHeap))-1 > len((*m.leftHeap)) {
			fmt.Println("swop1")
			n := m.rightHeap.Pop()
			m.leftHeap.Push(n)
		} else if len((*m.rightHeap))+1 < len((*m.leftHeap)) {
			fmt.Println("swop2")
			n := m.leftHeap.Pop()
			m.rightHeap.Push(n)
		}

		// Then the current median is computed as the root of the larger heap, if they
		// differ in count, or the average of the roots of the two heaps, if they are the
		// same size.
		if len((*m.rightHeap)) > len((*m.leftHeap)) {
			m.Median = float64((*m.rightHeap)[0])
		} else if len((*m.rightHeap)) < len((*m.leftHeap)) {
			m.Median = float64((*m.leftHeap)[0])
		} else {
			m.Median = float64((*m.rightHeap)[0]+(*m.leftHeap)[0]) / float64(2.0)
		}
	}
	m.Size += 1
	fmt.Println(*m.leftHeap, *m.rightHeap)
	if m.Size >= 2 {
		fmt.Println((*m.leftHeap)[0], (*m.rightHeap)[0])
	}

	return m.Median
}
