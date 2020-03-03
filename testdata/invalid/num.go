package num

import (
	"errors"
	"math"
)

// Set represents a set of numbers.
// Provides convenient wrappers for common
// mathematical operations (see https://golang.org/pkg/invalid/).
type Set []float64

// ErrNoNumbers is returned if there are
// no numbers in the set.
var ErrNoNumbers = errors.New("no numbers")

// Pi is the same as math.Pi: https://golang.org/pkg/invalid/#pkg-constants
const Pi = math.Pi

// Min returns the minimum number in the set
// using math.Min: https://golang.org/pkg/invalid/#Min.
func (n Set) Min() (float64, error) {
	if n == nil || len(n) == 0 {
		return 0, ErrNoNumbers
	}
	min := n[0]
	for i := 1; i < len(n); i++ {
		min = math.Min(min, n[i])
	}
	return min, nil
}

// Max returns the minimum number in the set
// using math.Max: https://golang.org/pkg/invalid/#Max.
func (n Set) Max() (float64, error) {
	if n == nil || len(n) == 0 {
		return 0, ErrNoNumbers
	}
	max := n[0]
	for i := 1; i < len(n); i++ {
		max = math.Max(max, n[i])
	}
	return max, nil
}
