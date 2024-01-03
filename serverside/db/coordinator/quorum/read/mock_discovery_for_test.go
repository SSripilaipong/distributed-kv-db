package read

import (
	"context"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/serverside/db/coordinator/quorum"
)

type discoveryMock[Key, Data any] struct {
	nodes func(ctx context.Context) rslt.Of[[]quorum.Node[Key, Data]]
}

func (m discoveryMock[Key, Data]) Nodes(ctx context.Context) rslt.Of[[]quorum.Node[Key, Data]] {
	return m.nodes(ctx)
}

var _ quorum.Discovery[int, int] = &discoveryMock[int, int]{}

func discovery[Key, Data any]() quorum.Discovery[Key, Data] {
	return discoveryWithNodes(rslt.Value([]quorum.Node[Key, Data]{}))
}

func discoveryWithNodes[Key, Data any](nodes rslt.Of[[]quorum.Node[Key, Data]]) quorum.Discovery[Key, Data] {
	return discoveryMock[Key, Data]{nodes: func(ctx context.Context) rslt.Of[[]quorum.Node[Key, Data]] {
		return nodes
	}}
}

func discoveryCaptureCtx[Key, Data any](ctx *context.Context) quorum.Discovery[Key, Data] {
	return discoveryMock[Key, Data]{nodes: func(c context.Context) rslt.Of[[]quorum.Node[Key, Data]] {
		*ctx = c
		return rslt.Value([]quorum.Node[Key, Data]{})
	}}
}
