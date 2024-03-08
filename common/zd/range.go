package zd

func RangeCh(start int, endEx int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := start; i < endEx; i++ {
			ch <- i
		}
	}()
	return ch
}
