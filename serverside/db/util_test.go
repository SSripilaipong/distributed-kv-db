package db

func interruptedSignal() <-chan struct{} {
	interrupt := make(chan struct{}, 1)
	interrupt <- struct{}{}
	return interrupt
}

func uninterruptedSignal() <-chan struct{} {
	return make(chan struct{})
}
