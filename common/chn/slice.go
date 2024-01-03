package chn

func NewFromSlice[T any](xs []T) chan T {
	ch := make(chan T, len(xs))
	for _, x := range xs {
		ch <- x
	}
	return ch
}
