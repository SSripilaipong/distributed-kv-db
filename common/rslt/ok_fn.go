package rslt

func OkFunc[A, B any](f func(A) Of[B]) func(A) bool {
	return func(x A) bool {
		return f(x).IsOk()
	}
}
