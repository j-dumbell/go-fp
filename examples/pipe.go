package examples

import (
	"github.com/j-dumbell/go-fp"
	"strconv"
)

func add3(i int) int {
	return i + 3
}

// "8"
var piped = fp.Pipe2(5, add3, strconv.Itoa)
