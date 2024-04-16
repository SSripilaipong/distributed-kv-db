package tuples

type Of4[T1, T2, T3, T4 any] struct {
	x1 T1
	x2 T2
	x3 T3
	x4 T4
}

func New4[T1, T2, T3, T4 any](x1 T1, x2 T2, x3 T3, x4 T4) Of4[T1, T2, T3, T4] {
	return Of4[T1, T2, T3, T4]{
		x1: x1,
		x2: x2,
		x3: x3,
		x4: x4,
	}
}

func Fn4[T1, T2, T3, T4, R any](f func(x1 T1, x2 T2, x3 T3, x4 T4) R) func(Of4[T1, T2, T3, T4]) R {
	return func(e Of4[T1, T2, T3, T4]) R {
		return f(e.x1, e.x2, e.x3, e.x4)
	}
}

func ExplodeFn4[T1, T2, T3, T4, R any](f func(Of4[T1, T2, T3, T4]) R) func(x1 T1, x2 T2, x3 T3, x4 T4) R {
	return func(x1 T1, x2 T2, x3 T3, x4 T4) R {
		return f(New4(x1, x2, x3, x4))
	}
}
