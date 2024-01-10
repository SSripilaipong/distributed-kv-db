package read

import (
	"context"
	"distributed-kv-db/common/chn"
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/common/zd"
	"distributed-kv-db/serverside/db/coordinator/quorum"
)

func New[Key any, Data orderableData, Node quorum.ReadNode[Key, Data]](discoverNodes quorum.DiscoverNodes[Key, Node]) quorum.ReadFunc[Key, Data] {
	return newFunc(
		discoverNodes,
		nil,
		latestData[Data],
	)
}

func newFunc[Key, Data, Node any](
	discoverNodes quorum.DiscoverNodes[Key, Node],
	readQuorumOfNodesData func(context.Context, Key, []Node) rslt.Of[[]Data],
	latestData func([]Data) rslt.Of[Data],
) quorum.ReadFunc[Key, Data] {
	return func(ctx context.Context, key Key) rslt.Of[Data] {
		readQuorumData := fn.Bind2(ctx, key, readQuorumOfNodesData)

		return fn.Compose3(
			rslt.FmapPartial(latestData), rslt.FmapPartial(readQuorumData), fn.Ctx(ctx, discoverNodes),
		)(key)
	}
}

func latestDataFromQuorumOfNodesFunc[Key, Data, Node any](
	readNodesDataToChannel func(context.Context, Key, []Node) <-chan Data,
	latestData func([]Data) rslt.Of[Data],
) func(context.Context, Key, []Node) rslt.Of[Data] {
	quorumOfData := fn.Compose(chn.FirstNFunc[Data], numberOfQuorum)

	return func(ctx context.Context, key Key, nodes []Node) rslt.Of[Data] {
		return fn.Compose3(
			rslt.FmapPartial(latestData), quorumOfData(len(nodes)), fn.Bind2(ctx, key, readNodesDataToChannel),
		)(nodes)
	}
}

var numberOfQuorum = fn.Compose(zd.Successor, zd.Half)
