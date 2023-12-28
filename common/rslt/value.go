package rslt

func Value[T any](v T) Of[T] {
	return Of[T]{
		value: v,
	}
}
