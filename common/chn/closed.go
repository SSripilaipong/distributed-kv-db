package chn

func Closed[T any](ch chan T) chan T {
	close(ch)
	return ch
}
