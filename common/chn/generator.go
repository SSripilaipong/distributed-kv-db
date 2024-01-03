package chn

func Repeat[T any](x T) <-chan T {
	ch := make(chan T)
	go func() {
		for {
			ch <- x
		}
	}()
	return ch
}
