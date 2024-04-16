package mergeread

import (
	"context"
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/common/slc"
	"distributed-kv-db/common/tuples"
	peerRead "distributed-kv-db/serverside/db/coordinator/peer/read"
	quorumRead "distributed-kv-db/serverside/db/coordinator/quorum/read"
)

func FromNodes[Key, Data any, Node peerRead.ReadableNode[Key, Data]](
	nReplicas uint,
	merge func(x, y Data) Data,
) func(ctx context.Context, node []Node, key Key) rslt.Of[Data] {
	return composeFromNodes[Key, Data, Node](
		merge, quorumRead.FromNodes[Key, Data, Node](nReplicas),
	)
}

func composeFromNodes[Key, Data, Node any](
	merge func(x, y Data) Data,
	qRead func(ctx context.Context, nodes []Node, key Key) rslt.Of[[]Data],
) func(ctx context.Context, nodes []Node, key Key) rslt.Of[Data] {

	return tuples.ExplodeFn3(fn.Compose(
		rslt.FmapPartial(slc.ReduceFn[[]Data](merge)), tuples.Fn3(qRead),
	))
}
