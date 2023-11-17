package examples

import (
	"github.com/j-dumbell/go-fp"
	"strconv"
)

func double(i int) int {
	return 2 * i
}

// func(int) string
var doubleThenString = fp.Flow2(double, strconv.Itoa)

// "4"
var flowed = doubleThenString(2)
