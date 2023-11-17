package fp

func Flow2[A, B, C any](ab func(A) B, bc func(B) C) func(A) C {
	return func(a A) C {
		return bc(ab(a))
	}
}

func Flow3[A, B, C, D any](ab func(A) B, bc func(B) C, cd func(C) D) func(A) D {
	return func(a A) D {
		return cd(bc(ab(a)))
	}
}

func Flow4[A, B, C, D, E any](ab func(A) B, bc func(B) C, cd func(C) D, de func(D) E) func(A) E {
	return func(a A) E {
		return de(cd(bc(ab(a))))
	}
}

func Flow5[A, B, C, D, E, F any](ab func(A) B, bc func(B) C, cd func(C) D, de func(D) E, ef func(E) F) func(A) F {
	return func(a A) F {
		return ef(de(cd(bc(ab(a)))))
	}
}
