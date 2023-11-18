package fp

// MapI is a curried array map implementation with mandatory index parameter
// on the callback fn.
func MapI[A, B any](fn func(A, int) B) func([]A) []B {
	return func(arr []A) []B {
		mapped := make([]B, len(arr))
		for i, a := range arr {
			mapped[i] = fn(a, i)
		}
		return mapped
	}
}

// Map is a curried array map implementation.
func Map[A, B any](fn func(A) B) func([]A) []B {
	fnI := func(a A, i int) B {
		return fn(a)
	}

	return MapI(fnI)
}

// FilterI is a curried array filter implementation with mandatory index parameter
// on the callback.
func FilterI[A any](fn func(A, int) bool) func([]A) []A {
	return func(arr []A) []A {
		var filtered []A
		for i, a := range arr {
			if fn(a, i) {
				filtered = append(filtered, a)
			}
		}
		return filtered
	}
}

// Filter is a curried array filter implementation.
func Filter[A any](fn func(A) bool) func([]A) []A {
	fnI := func(a A, i int) bool {
		return fn(a)
	}
	return FilterI(fnI)
}

// FlatMap creates a new slice by transforming and flattening another slice by calling the
// provided callback on each element.
func FlatMap[A, B any](fn func(A) []B) func([]A) []B {
	return func(arr []A) []B {
		var mapped []B
		for _, a := range arr {
			mapped = append(mapped, fn(a)...)
		}
		return mapped
	}
}

// FlatMapI creates a new slice by transforming and flattening another slice by calling the
// provided callback on each element.  The callback takes a 2nd argument for the slice index.
func FlatMapI[A, B any](fn func(A, int) []B) func([]A) []B {
	return func(arr []A) []B {
		var mapped []B
		for i, a := range arr {
			mapped = append(mapped, fn(a, i)...)
		}
		return mapped
	}
}
