package rslt

func Error[T any](err error) Of[T] {
	return Of[T]{
		isError: true,
		err:     err,
	}
}

func OfError(r Of[error]) error {
	if r.IsOk() {
		return r.Value()
	}
	return r.Error()
}

func ResultOrError[T any](r Of[T], err error) Of[T] {
	if r.IsError() || err == nil {
		return r
	}
	return Error[T](err)
}
