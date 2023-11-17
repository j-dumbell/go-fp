package fp

// Curry2 curries the provided arity 2 function from left to right.
func Curry2[A, B, C any](fn func(A, B) C) func(A) func(B) C {
	return func(a A) func(B) C {
		return func(b B) C {
			return fn(a, b)
		}
	}
}

// Curry3 curries the provided arity 3 function from left to right.
func Curry3[A, B, C, D any](fn func(A, B, C) D) func(A) func(B) func(C) D {
	return func(a A) func(B) func(C) D {
		return func(b B) func(C) D {
			return func(c C) D {
				return fn(a, b, c)
			}
		}
	}
}

// Curry4 curries the provided arity 4 function from left to right.
func Curry4[A, B, C, D, E any](fn func(A, B, C, D) E) func(A) func(B) func(C) func(D) E {
	return func(a A) func(B) func(C) func(D) E {
		return func(b B) func(C) func(D) E {
			return func(c C) func(D) E {
				return func(d D) E {
					return fn(a, b, c, d)
				}
			}
		}
	}
}

// CurryR2 curries the provided arity 2 function from right to left.
func CurryR2[A, B, C any](fn func(A, B) C) func(B) func(A) C {
	return func(b B) func(A) C {
		return func(a A) C {
			return fn(a, b)
		}
	}
}

// CurryR3 curries the provided arity 3 function from right to left.
func CurryR3[A, B, C, D any](fn func(A, B, C) D) func(C) func(B) func(A) D {
	return func(c C) func(B) func(A) D {
		return func(b B) func(A) D {
			return func(a A) D {
				return fn(a, b, c)
			}
		}
	}
}

// CurryR4 curries the provided arity 4 function from right to left.
func CurryR4[A, B, C, D, E any](fn func(A, B, C, D) E) func(D) func(C) func(B) func(A) E {
	return func(d D) func(C) func(B) func(A) E {
		return func(c C) func(B) func(A) E {
			return func(b B) func(A) E {
				return func(a A) E {
					return fn(a, b, c, d)
				}
			}
		}
	}
}
