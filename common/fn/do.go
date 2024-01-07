package fn

func Do[A, B any](f func(A) B) func(A) {
	return func(x A) {
		_ = f(x)
	}
}
