package read

import (
	"context"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/serverside/db/coordinator/quorum"
)

type discoveryMock[Key, Data any] struct {
	nodes func(ctx context.Context, key Key) rslt.Of[[]quorum.Node[Key, Data]]
}

func (m discoveryMock[Key, Data]) Nodes(ctx context.Context, key Key) rslt.Of[[]quorum.Node[Key, Data]] {
	return m.nodes(ctx, key)
}

var _ quorum.Discovery[int, int] = &discoveryMock[int, int]{}

func nodesFuncWithResult[Key, Data any](nodes rslt.Of[[]quorum.Node[Key, Data]]) func(ctx context.Context, key Key) rslt.Of[[]quorum.Node[Key, Data]] {
	return func(ctx context.Context, key Key) rslt.Of[[]quorum.Node[Key, Data]] {
		return nodes
	}
}

func nodesFuncCaptureContext[Key, Data any](ctx *context.Context) func(ctx context.Context, key Key) rslt.Of[[]quorum.Node[Key, Data]] {
	return func(c context.Context, key Key) rslt.Of[[]quorum.Node[Key, Data]] {
		*ctx = c
		return rslt.Value([]quorum.Node[Key, Data]{})
	}
}

func nodesFuncCaptureKey[Key, Data any](key *Key) func(c context.Context, k Key) rslt.Of[[]quorum.Node[Key, Data]] {
	return func(c context.Context, k Key) rslt.Of[[]quorum.Node[Key, Data]] {
		*key = k
		return rslt.Value([]quorum.Node[Key, Data]{})
	}
}

func discoveryWithNodesFunc[Key, Data any](f func(ctx context.Context, key Key) rslt.Of[[]quorum.Node[Key, Data]]) quorum.Discovery[Key, Data] {
	return discoveryMock[Key, Data]{nodes: f}
}
