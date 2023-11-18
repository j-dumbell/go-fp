package examples

import "github.com/j-dumbell/go-fp"

var nums = []int{3, -1, -4}

// []int{3, -1}
var positiveNumbers = fp.FilterI(func(i, j int) bool {
	return i+j >= 0
})(nums)
