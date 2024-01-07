package grt

func Do[T any](f func(T)) func(T) {
	return func(x T) {
		go f(x)
	}
}
