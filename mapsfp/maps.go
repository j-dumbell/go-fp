package mapsfp

import (
	"maps"
)

// Equal is a curried version of maps.Equal.
func Equal[M1, M2 ~map[K]V, K, V comparable](m1 M1) func(m2 M2) bool {
	return func(m2 M2) bool {
		return maps.Equal(m1, m2)
	}
}

// EqualFunc is a curried version of maps.EqualFunc.
func EqualFunc[M1 ~map[K]V1, M2 ~map[K]V2, K comparable, V1, V2 any](eq func(V1, V2) bool) func(m1 M1) func(m2 M2) bool {
	return func(m1 M1) func(M2) bool {
		return func(m2 M2) bool {
			return maps.EqualFunc(m1, m2, eq)
		}
	}
}
