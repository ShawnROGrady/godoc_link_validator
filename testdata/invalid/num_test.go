package num

import "fmt"

// Similar to math.Min: https://golang.org/invalid/invalid/#Min.
func ExampleSet_Min() {
	s := Set([]float64{10, 9, -1, 3, 4})
	min := s.Min()
	fmt.Printf("%v\n", min)
}
