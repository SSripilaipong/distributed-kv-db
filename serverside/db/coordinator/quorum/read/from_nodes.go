package read

import (
	"context"
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/common/tuples"
	peerRead "distributed-kv-db/serverside/db/coordinator/peer/read"
	quorumFilter "distributed-kv-db/serverside/db/coordinator/quorum/filter"
)

func FromNodes[Key, Data any, Node peerRead.ReadableNode[Key, Data]](
	nReplicas uint,
) func(ctx context.Context, nodes []Node, key Key) rslt.Of[[]Data] {
	return composeFromNodes[Key, Data, Node](
		quorumFilter.ChannelToSlice[Data](nReplicas),
		peerRead.NodesDataToChannel[Key, Data, Node],
	)
}

func composeFromNodes[Key, Data, Node any](
	qFilter func(ch <-chan Data) rslt.Of[[]Data],
	pRead func(ctx context.Context, key Key, nodes []Node) <-chan Data,
) func(ctx context.Context, nodes []Node, key Key) rslt.Of[[]Data] {
	return tuples.ExplodeFn3(fn.Compose(
		qFilter, tuples.Fn3(swapArg2And3(pRead)),
	))
}

func swapArg2And3[A, B, C, R any](f func(A, B, C) R) func(A, C, B) R {
	return func(a A, c C, b B) R {
		return f(a, b, c)
	}
}
