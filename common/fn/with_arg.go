package fn

func WithArg[A1, A2, B any](x1 A1, f func(A1, A2) B) func(A2) B {
	return func(x2 A2) B {
		return f(x1, x2)
	}
}

func WithArg2[A1, A2, A3, B any](x1 A1, x2 A2, f func(A1, A2, A3) B) func(A3) B {
	return func(x3 A3) B {
		return f(x1, x2, x3)
	}
}

func WithArg3[A1, A2, A3, A4, B any](x1 A1, x2 A2, x3 A3, f func(A1, A2, A3, A4) B) func(A4) B {
	return func(x4 A4) B {
		return f(x1, x2, x3, x4)
	}
}
