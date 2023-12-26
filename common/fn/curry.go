package fn

func Uncurry[A1, A2, B any](f func(A1) func(A2) B) func(A1, A2) B {
	return func(x1 A1, x2 A2) B {
		return f(x1)(x2)
	}
}
