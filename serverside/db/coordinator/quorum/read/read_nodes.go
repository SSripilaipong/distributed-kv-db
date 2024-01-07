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

func readNodesDataToChannel[Key, Data any, Node quorum.ReadNode[Key, Data]](ctx context.Context, key Key, nodes []Node) <-chan Data {
	var wg sync.WaitGroup
	wg.Add(len(nodes))
	ch := make(chan Data, len(nodes))

	slc.Do(grt.Do(readNodeDataToChannel[Key, Data, Node](ctx, &wg, ch, key)), nodes)
	go chn.CloseAfterWaitGroup(&wg, ch)

	return ch
}

func readNodeDataToChannel[Key, Data any, Node quorum.ReadNode[Key, Data]](ctx context.Context, wg *sync.WaitGroup, ch chan<- Data, key Key) func(node Node) {
	return fn.Do(wgrp.MustDone(wg, fn.Compose(
		rslt.Fmap(chn.SendToWithContext(ctx, ch)), readNode[Data, Key, Node](ctx, key),
	)))
}

func readNode[Data, Key any, Node quorum.ReadNode[Key, Data]](ctx context.Context, key Key) func(node Node) rslt.Of[Data] {
	return func(node Node) rslt.Of[Data] {
		return node.Read(ctx, key)
	}
}
