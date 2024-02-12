package rslt

func Value[T any](v T) Of[T] {
	return Of[T]{
		value: v,
	}
}

func ValueOf[T any](r Of[T]) T {
	return r.Value()
}
