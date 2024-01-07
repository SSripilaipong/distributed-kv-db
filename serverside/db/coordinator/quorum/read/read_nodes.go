package read

import (
	"context"
	"distributed-kv-db/common/chn"
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/grt"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/common/slc"
	"distributed-kv-db/common/wgrp"
	"distributed-kv-db/serverside/db/coordinator/quorum"
	"sync"
)

func readNodesDataToChannel[Key, Data any](ctx context.Context, key Key, nodes []quorum.ReadableNode[Key, Data]) <-chan Data {
	var wg sync.WaitGroup
	wg.Add(len(nodes))
	ch := make(chan Data, len(nodes))

	slc.Do(grt.Do(readNodeDataToChannel(ctx, &wg, ch, key)), nodes)
	go chn.CloseAfterWaitGroup(&wg, ch)

	return ch
}

func readNodeDataToChannel[Key, Data any](ctx context.Context, wg *sync.WaitGroup, ch chan<- Data, key Key) func(node quorum.ReadableNode[Key, Data]) {
	return fn.Do(wgrp.MustDone(wg, fn.Compose(
		rslt.Fmap(chn.SendToWithContext(ctx, ch)), readNode[Data](ctx, key),
	)))
}

func readNode[Data, Key any](ctx context.Context, key Key) func(node quorum.ReadableNode[Key, Data]) rslt.Of[Data] {
	return func(node quorum.ReadableNode[Key, Data]) rslt.Of[Data] {
		return node.Read(ctx, key)
	}
}
