package readrepair

import (
	"context"
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/rslt"
	peerRead "distributed-kv-db/serverside/db/coordinator/peer/read"
	quorumMergeRead "distributed-kv-db/serverside/db/coordinator/quorum/mergeread"
)

func FromNodes[Key, Data any, Node peerRead.ReadableNode[Key, Data]](
	nReplicas uint,
	merge func(x Data, y Data) Data,
) func(context.Context, []Node, Key) rslt.Of[Data] {
	return composeFromNodes[Key, Data, Node](
		quorumMergeRead.FromNodes[Key, Data, Node](nReplicas, merge),
		nil, // TODO inject real functions
	)
}

func composeFromNodes[Key, Data, Node any](
	qMergeRead func(ctx context.Context, nodes []Node, key Key) rslt.Of[Data],
	qBlindWrite func(ctx context.Context, nodes []Node, key Key, data Data) error,
) func(context.Context, []Node, Key) rslt.Of[Data] {
	return func(ctx context.Context, nodes []Node, key Key) rslt.Of[Data] {
		writeFn := fmapResultToError(fn.WithArg3(ctx, nodes, key, qBlindWrite))

		result := qMergeRead(ctx, nodes, key)
		return rslt.ResultOrError(result, writeFn(result))
	}
}
