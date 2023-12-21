package result

func Error[T any](err error) Of[T] {
	return Of[T]{
		isError: true,
		err:     err,
	}
}
