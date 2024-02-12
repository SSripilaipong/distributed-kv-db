package read

import (
	"context"
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/common/typ"
	"distributed-kv-db/serverside/db/coordinator/peer/discovery"
	peerRead "distributed-kv-db/serverside/db/coordinator/peer/read"
	quorumFilter "distributed-kv-db/serverside/db/coordinator/quorum/filter"
)

func NodesToDataSlice[Key, Data any, Node peerRead.ReadableNode[Key, Data]](
	discoverNodes discovery.Func[Key, Node],
) func(ctx context.Context, key Key) rslt.Of[[]Data] {
	return composeNodesToDataSlice[Key, Data, Node](
		quorumFilter.ChannelToSlice[Data], peerRead.NodesDataToChannel[Key, Data, Node], discoverNodes,
	)
}

func composeNodesToDataSlice[Key, Data, Node any](
	filterQuorum func(n int) func(<-chan Data) rslt.Of[[]Data],
	readNodes func(ctx context.Context, key Key, nodes []Node) <-chan Data,
	discoverNodes discovery.Func[Key, Node],
) func(ctx context.Context, key Key) rslt.Of[[]Data] {
	filterQuorumForNodes := fn.Compose3(filterQuorum, rslt.ValueOf[int], rslt.Fmap(lenOfSlice[Node]))

	return func(ctx context.Context, key Key) rslt.Of[[]Data] {
		nodes := discoverNodes(ctx, key)
		quorumOfN := filterQuorumForNodes(nodes)
		rslt.Fmap(fn.Compose(quorumOfN, fn.WithArg2(ctx, key, readNodes)))(nodes)
		return rslt.Value(typ.Zero[[]Data]())
	}
}

func lenOfSlice[T any](xs []T) int {
	return len(xs)
}
