package fn

func Bind[A1, A2, B any](x1 A1, f func(A1, A2) B) func(A2) B {
	return func(x2 A2) B {
		return f(x1, x2)
	}
}

func Bind2[A1, A2, A3, B any](x1 A1, x2 A2, f func(A1, A2, A3) B) func(A3) B {
	return func(x3 A3) B {
		return f(x1, x2, x3)
	}
}
