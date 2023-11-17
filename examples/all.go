package examples

import (
	fp "github.com/j-dumbell/go-fp"
	"strconv"
)

var ints = []int{1, 2, 3}

func square(i int) int {
	return i * i
}

func sum(ints []int) int {
	total := 0
	for _, i := range ints {
		total += i
	}
	return total
}

// "14"
var res = fp.Pipe3(ints, fp.Map(square), sum, strconv.Itoa)
