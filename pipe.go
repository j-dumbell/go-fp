package fp

func Pipe2[A, B, C any](a A, ab func(A) B, bc func(B) C) C {
	return Flow2(ab, bc)(a)
}

func Pipe3[A, B, C, D any](a A, ab func(A) B, bc func(B) C, cd func(C) D) D {
	return Flow3(ab, bc, cd)(a)
}

func Pipe4[A, B, C, D, E any](a A, ab func(A) B, bc func(B) C, cd func(C) D, de func(D) E) E {
	return Flow4(ab, bc, cd, de)(a)
}

func Pipe5[A, B, C, D, E, F any](a A, ab func(A) B, bc func(B) C, cd func(C) D, de func(D) E, ef func(E) F) F {
	return Flow5(ab, bc, cd, de, ef)(a)
}
