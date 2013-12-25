package main

import (
	"fmt"
	"github.com/hermanschaaf/algorithms/median"
)

func main() {
	lst := []int{1, 2, 3, 4, 5}
	sm := median.StreamingMedian{}
	for i := range lst {
		m := sm.Add(lst[i])
		fmt.Printf("Added %v, new median is %v\n", lst[i], m)
	}
}
