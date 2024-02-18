package wgrp

import "sync"

func MustDone[A, B any](wg *sync.WaitGroup, f func(A) B) func(A) B {
	return func(x A) B {
		defer wg.Done()
		return f(x)
	}
}

func MustDoneAfterDo[A any](wg *sync.WaitGroup, f func(A)) func(A) {
	return func(x A) {
		defer wg.Done()
		f(x)
	}
}
