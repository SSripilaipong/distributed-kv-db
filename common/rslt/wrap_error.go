package rslt

import "fmt"

func WrapErrorFunc[T any](s string) func(x Of[T]) Of[T] {
	return func(x Of[T]) Of[T] {
		if x.IsOk() {
			return x
		}
		return Error[T](fmt.Errorf("%s: %w", s, x.Error()))
	}
}
