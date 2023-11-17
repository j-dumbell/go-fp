package examples

import (
	"github.com/j-dumbell/go-fp"
	"strconv"
)

var numbers = []int{1, 2, 3}

// []string{"1", "2", "3"}
var toStrings = fp.Map(strconv.Itoa)(numbers)
