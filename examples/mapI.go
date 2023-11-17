package examples

import "github.com/j-dumbell/go-fp"

var uints = []int{1, 2, 3}

// []string{1, 3, 5}
var result = fp.MapI(func(i int, index int) int {
	return i + index
})(uints)
