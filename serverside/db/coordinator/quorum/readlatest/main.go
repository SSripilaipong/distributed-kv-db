package readlatest

import (
	"context"
	"distributed-kv-db/common/chn"
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/common/zd"
	"distributed-kv-db/serverside/db/coordinator/peer/discovery"
	peerRead "distributed-kv-db/serverside/db/coordinator/peer/read"
	quorumRead "distributed-kv-db/serverside/db/coordinator/quorum/read"
)

func New[Key any, Data ReadableData, Node peerRead.ReadableNode[Key, Data]](discoverNodes discovery.Func[Key, Node]) quorumRead.Func[Key, Data] {
	return newFunc(
		discoverNodes,
		nil,
		latestData[Data],
	)
}

func newFunc[Key, Data, Node any](
	discoverNodes discovery.Func[Key, Node],
	readQuorumOfNodesData func(context.Context, Key, []Node) rslt.Of[[]Data],
	latestData func([]Data) rslt.Of[Data],
) quorumRead.Func[Key, Data] {
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
