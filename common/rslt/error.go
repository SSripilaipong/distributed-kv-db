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

func ValueOrError[T any](x T, err error) Of[T] {
	if err != nil {
		return Error[T](err)
	}
	return Value(x)
}
