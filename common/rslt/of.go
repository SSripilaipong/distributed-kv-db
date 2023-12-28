package rslt

type Of[T any] struct {
	isError bool
	value   T
	err     error
}

func (r Of[T]) Error() error {
	return r.err
}

func (r Of[T]) Value() T {
	return r.value
}

func (r Of[T]) IsOk() bool {
	return !r.isError
}

func (r Of[T]) IsError() bool {
	return r.isError
}
