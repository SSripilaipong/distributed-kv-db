package strm

import (
	"context"
	"distributed-kv-db/common/chn"
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/prll"
	"distributed-kv-db/common/rslt"
	"sync"
)

func MapResultFromSlice[A, B any](f func(A) rslt.Of[B], ctx context.Context, xs []A) <-chan B {
	var wg sync.WaitGroup
	ch := make(chan B, len(xs))

	prll.ApplySlice(&wg, fn.Do(fn.Compose(rslt.Fmap(chn.SendToWithContext(ctx, ch)), f)), xs)
	go chn.CloseAfterWaitGroup(&wg, ch)

	return ch
}

func MapSlice[A, B any](f func(A) B, ctx context.Context, xs []A) <-chan B {
	var wg sync.WaitGroup
	ch := make(chan B, len(xs))

	prll.ApplySlice(&wg, fn.Do(fn.Compose(chn.SendToWithContext(ctx, ch), f)), xs)
	go chn.CloseAfterWaitGroup(&wg, ch)

	return ch
}
