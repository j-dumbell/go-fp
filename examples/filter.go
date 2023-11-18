package examples

import "github.com/j-dumbell/go-fp"

func isHello(s string) bool {
	return s == "hello"
}

var words = []string{"foo", "hello", "world"}

// []int{"hello"}
var onlyHello = fp.Filter(isHello)(words)
