package rslt

func Resolve[T, R any](f func(T) R, g func(error) R, x Of[T]) R {
	if x.IsOk() {
		return f(x.Value())
	}
	return g(x.Error())
}

func ResolveFunc[T, R any](f func(T) R, g func(error) R) func(Of[T]) R {
	return func(x Of[T]) R {
		return Resolve(f, g, x)
	}
}
