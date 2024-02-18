package slc

func Map[S ~[]A, A, B any](f func(A) B, xs S) (ys []B) {
	for _, x := range xs {
		ys = append(ys, f(x))
	}
	return
}
