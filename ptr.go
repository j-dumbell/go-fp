package fp

// ToPtr returns a pointer to t.
func ToPtr[T any](t T) *T {
	return &t
}

// FromPtr returns the pointer's value.
func FromPtr[T any](t *T) T {
	return *t
}
