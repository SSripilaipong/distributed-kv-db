package chn

import "sync"

func CloseAfterWaitGroup[T any](wg *sync.WaitGroup, ch chan<- T) {
	wg.Wait()
	close(ch)
}
