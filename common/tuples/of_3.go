package tuples

type Of3[T1, T2, T3 any] struct {
	x1 T1
	x2 T2
	x3 T3
}

func New3[T1, T2, T3 any](x1 T1, x2 T2, x3 T3) Of3[T1, T2, T3] {
	return Of3[T1, T2, T3]{
		x1: x1,
		x2: x2,
		x3: x3,
	}
}

func Fn3[T1, T2, T3, R any](f func(x1 T1, x2 T2, x3 T3) R) func(Of3[T1, T2, T3]) R {
	return func(e Of3[T1, T2, T3]) R {
		return f(e.x1, e.x2, e.x3)
	}
}

func ExplodeFn3[T1, T2, T3, R any](f func(Of3[T1, T2, T3]) R) func(x1 T1, x2 T2, x3 T3) R {
	return func(x1 T1, x2 T2, x3 T3) R {
		return f(New3(x1, x2, x3))
	}
}
