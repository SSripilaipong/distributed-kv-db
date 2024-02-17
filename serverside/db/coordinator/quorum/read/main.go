package read

import (
	"context"
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/serverside/db/coordinator/peer/discovery"
	peerRead "distributed-kv-db/serverside/db/coordinator/peer/read"
	quorumFilter "distributed-kv-db/serverside/db/coordinator/quorum/filter"
)

func NodesToDataSlice[Key, Data any, Node peerRead.ReadableNode[Key, Data]](
	nReplicas int, discoverNodes discovery.Func[Key, Node],
) func(ctx context.Context, key Key) rslt.Of[[]Data] {
	return composeNodesToDataSlice[Key, Data, Node](
		quorumFilter.ChannelToSlice[Data](nReplicas), peerRead.NodesDataToChannel[Key, Data, Node], discoverNodes,
	)
}

func composeNodesToDataSlice[Key, Data, Node any](
	filterQuorum func(<-chan Data) rslt.Of[[]Data],
	readNodes func(ctx context.Context, key Key, nodes []Node) <-chan Data,
	discoverNodes discovery.Func[Key, Node],
) func(ctx context.Context, key Key) rslt.Of[[]Data] {
	return func(ctx context.Context, key Key) rslt.Of[[]Data] {
		filterAfterRead := fn.Compose(filterQuorum, fn.WithArg2(ctx, key, readNodes))
		return fn.Compose(
			rslt.FmapPartial(filterAfterRead), fn.Ctx(ctx, discoverNodes),
		)(key)
	}
}
