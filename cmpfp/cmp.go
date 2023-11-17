package cmpfp

import (
	"cmp"
)

// Compare is a curried version of cmp.Compare.
func Compare[T cmp.Ordered](x T) func(y T) int {
	return func(y T) int {
		return cmp.Compare(x, y)
	}
}

// Less is a curried version of cmp.Less.
func Less[T cmp.Ordered](x T) func(y T) bool {
	return func(y T) bool {
		return cmp.Less(x, y)
	}
}
