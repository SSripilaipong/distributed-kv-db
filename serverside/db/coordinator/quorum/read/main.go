package read

import (
	"context"
	"distributed-kv-db/common/chn"
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/common/zd"
	"distributed-kv-db/serverside/db/coordinator/quorum"
)

func New[Key, Data any](discovery quorum.Discovery[Key, Data]) quorum.ReadFunc[Key, Data] {
	return newFunc(
		discovery,
		nil, // TODO inject this
		nil, // TODO inject this
	)
}

func newFunc[Key, Data any](
	discovery quorum.Discovery[Key, Data],
	readNodeDataToChannel func([]quorum.Node[Key, Data]) <-chan Data,
	latestData func([]Data) Data,
) quorum.ReadFunc[Key, Data] {
	latestDataFromQuorum := latestDataFromQuorumOfNodesFunc[Key, Data](readNodeDataToChannel, latestData)

	return fn.Uncurry(func(ctx context.Context) func(Key) rslt.Of[Data] {
		return fn.Compose(
			rslt.FmapPartial(fn.Ctx(ctx, latestDataFromQuorum)), fn.Ctx(ctx, discovery.Nodes),
		)
	})
}

func latestDataFromQuorumOfNodesFunc[Key, Data any](
	readNodeDataToChannel func([]quorum.Node[Key, Data]) <-chan Data,
	latestData func([]Data) Data,
) func(ctx context.Context, nodes []quorum.Node[Key, Data]) rslt.Of[Data] {
	quorumOfData := fn.Compose(chn.FirstNFunc[Data], numberOfQuorum)

	return func(ctx context.Context, nodes []quorum.Node[Key, Data]) rslt.Of[Data] {
		latestDataFromQuorum := fn.Compose3(
			rslt.Fmap(latestData), quorumOfData(len(nodes)), readNodeDataToChannel,
		)
		return latestDataFromQuorum(nodes)
	}
}

var numberOfQuorum = fn.Compose(zd.Successor, zd.Half)
