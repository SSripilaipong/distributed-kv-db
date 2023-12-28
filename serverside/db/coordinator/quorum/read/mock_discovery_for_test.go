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

func discoveryWithNodes[Key, Data any](nodes rslt.Of[[]quorum.Node[Key, Data]]) quorum.Discovery[Key, Data] {
	return discoveryMock[Key, Data]{nodes: func(ctx context.Context) rslt.Of[[]quorum.Node[Key, Data]] {
		return nodes
	}}
}
