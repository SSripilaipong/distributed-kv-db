package strm

import (
	"context"
	"distributed-kv-db/common/chn"
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/grt"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/common/slc"
	"distributed-kv-db/common/wgrp"
	"sync"
)

func MapSlice[A, B any](f func(A) rslt.Of[B], ctx context.Context, xs []A) <-chan B {
	var wg sync.WaitGroup
	wg.Add(len(xs))
	ch := make(chan B, len(xs))

	slc.Do(grt.Do(executeToChannelWithWaitGroup(ctx, &wg, ch, f)), xs)
	go chn.CloseAfterWaitGroup(&wg, ch)

	return ch
}

func executeToChannelWithWaitGroup[A, B any](ctx context.Context, wg *sync.WaitGroup, ch chan<- B, f func(A) rslt.Of[B]) func(A) {
	return fn.Do(wgrp.MustDone(wg, fn.Compose(
		rslt.Fmap(chn.SendToWithContext(ctx, ch)), f),
	))
}
