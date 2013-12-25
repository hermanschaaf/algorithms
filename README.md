algorithms
==========

Common useful algorithms implemented in Go. Currently includes only a streaming median implementation.

Streaming Median
----------------

Streaming Median computes the running median of a list of values as they are added. This is useful when large datasets or different threads (goroutines) are involved, and sorting would be too costly. The algorithm is implemented internally using two heaps. It adds a number and computes the new median in O(logN), and can compute the median of the full list in O(NlogN), where N is the length of the list.

```
import (
    "fmt"
    "github.com/hermanschaaf/algorithms/median"
)

func main(){
   lst := []int{1,2,3,4,5}
   sm := median.StreamingMedian{}
   for i := range lst {
       m := sm.Add(lst[i])
       fmt.Printf("Added %v, new median is %v\n", lst[i], m)
   }
}
```

And the output is:

```
Added 1, new median is 1
Added 2, new median is 1.5
Added 3, new median is 2
Added 4, new median is 2.5
Added 5, new median is 3
```