package num

import (
	"fmt"
	"testing"
)

// Similar to math.Min: https://golang.org/invalid/invalid/#Min.
func ExampleSet_Min() {
	s := Set([]float64{10, 9, -1, 3, 4})
	min := s.Min()
	fmt.Printf("%v\n", min)
}

// Similar to math.Min: https://golang.org/invalid/invalid/#Min.
func TestMin(t *testing.T) {
	s := Set([]float64{10, 9, -1, 3, 4})
	min := s.Min()
	if min != -1 {
		t.Errorf("unexpected min (expected=%v, actual=%v)", -1, min)
	}
}
