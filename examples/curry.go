package examples

import "github.com/j-dumbell/go-fp"

func add(i, j int) int {
	return i + j
}

// func(int) func(int) int
var curriedAdd = fp.Curry2(add)

// 3
var oneAddTwo = curriedAdd(1)(2)
