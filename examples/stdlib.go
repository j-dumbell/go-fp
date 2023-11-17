package examples

import (
	"github.com/j-dumbell/go-fp/stringsfp"
)

var data = []string{"1", "2", "3"}

// Equivalent to strings.Join(data, ",")
var joinedData = stringsfp.Join(",")(data)
