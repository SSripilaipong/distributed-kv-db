package slc

func Map[A, B any](f func(A) B, xs []A) (ys []B) {
	for _, x := range xs {
		ys = append(ys, f(x))
	}
	return
}

func MapFunc[A, B any](f func(A) B) func(xs []A) (ys []B) {
	return func(xs []A) (ys []B) {
		return Map(f, xs)
	}
}
