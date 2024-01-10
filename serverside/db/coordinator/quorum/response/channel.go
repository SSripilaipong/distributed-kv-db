package response

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

func Channel[Response, Node any](ctx context.Context, request func(Node) rslt.Of[Response], nodes []Node) <-chan Response {
	var wg sync.WaitGroup
	wg.Add(len(nodes))
	ch := make(chan Response, len(nodes))

	slc.Do(grt.Do(requestToNode[Response, Node](ctx, &wg, ch, request)), nodes)
	go chn.CloseAfterWaitGroup(&wg, ch)

	return ch
}

func requestToNode[Response, Node any](ctx context.Context, wg *sync.WaitGroup, ch chan<- Response, request func(Node) rslt.Of[Response]) func(node Node) {
	return fn.Do(wgrp.MustDone(wg, fn.Compose(
		rslt.Fmap(chn.SendToWithContext(ctx, ch)), request),
	))
}
