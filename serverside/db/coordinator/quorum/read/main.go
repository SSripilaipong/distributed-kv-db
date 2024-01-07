package read

import (
	"context"
	"distributed-kv-db/common/chn"
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/common/zd"
	"distributed-kv-db/serverside/db/coordinator/quorum"
)

func New[Key, Data any, Node quorum.ReadNode[Key, Data]](discoverNodes quorum.DiscoverNodes[Key, Node]) quorum.ReadFunc[Key, Data] {
	return newFunc(
		discoverNodes,
		readNodesDataToChannel[Key, Data, Node],
		nil, // TODO inject this
	)
}

func newFunc[Key, Data, Node any](
	discoverNodes quorum.DiscoverNodes[Key, Node],
	readNodesDataToChannel func(context.Context, Key, []Node) <-chan Data,
	latestData func([]Data) Data,
) quorum.ReadFunc[Key, Data] {
	latestDataFromQuorum := latestDataFromQuorumOfNodesFunc(readNodesDataToChannel, latestData)

	return func(ctx context.Context, key Key) rslt.Of[Data] {
		return fn.Compose(
			rslt.FmapPartial(fn.Bind2(ctx, key, latestDataFromQuorum)), fn.Ctx(ctx, discoverNodes),
		)(key)
	}
}

func latestDataFromQuorumOfNodesFunc[Key, Data, Node any](
	readNodesDataToChannel func(context.Context, Key, []Node) <-chan Data,
	latestData func([]Data) Data,
) func(context.Context, Key, []Node) rslt.Of[Data] {
	quorumOfData := fn.Compose(chn.FirstNFunc[Data], numberOfQuorum)

	return func(ctx context.Context, key Key, nodes []Node) rslt.Of[Data] {
		return fn.Compose3(
			rslt.Fmap(latestData), quorumOfData(len(nodes)), fn.Bind2(ctx, key, readNodesDataToChannel),
		)(nodes)
	}
}

var numberOfQuorum = fn.Compose(zd.Successor, zd.Half)
