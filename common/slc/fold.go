package slc

func Fold[S ~[]E, A, E any](f func(A, E) A, z A, xs S) A {
	r := z
	for _, x := range xs {
		r = f(r, x)
	}
	return r
}
