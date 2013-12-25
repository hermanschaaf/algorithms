package median

import (
	"testing"
)

type Case struct {
	input  int
	median float64
}

func TestStreamingMedian(t *testing.T) {
	cases := []Case{
		Case{1, 1.0},
		Case{2, 1.5},
		Case{3, 2.0},
		Case{4, 2.5},
	}

	sm := StreamingMedian{}
	for k := range cases {
		a := sm.Add(cases[k].input)
		if a != cases[k].median {
			t.Errorf("Expected StreamingMedian.Add(%v) = %v, but got %v", cases[k].input, cases[k].median, a)
		}
	}
}
