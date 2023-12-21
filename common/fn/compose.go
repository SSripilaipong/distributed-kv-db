package fn

func Compose[A, B, C any](f func(B) C, g func(A) B) func(A) C {
	return func(x A) C {
		return f(g(x))
	}
}

func Compose3[A, B, C, D any](f func(C) D, g func(B) C, h func(A) B) func(A) D {
	return func(x A) D {
		return f(g(h(x)))
	}
}
