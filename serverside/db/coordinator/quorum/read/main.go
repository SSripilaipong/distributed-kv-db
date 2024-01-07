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
		readNodesDataToChannel[Key, Data],
		nil, // TODO inject this
	)
}

func newFunc[Key, Data any](
	discovery quorum.Discovery[Key, Data],
	readNodesDataToChannel func(context.Context, Key, []quorum.Node[Key, Data]) <-chan Data,
	latestData func([]Data) Data,
) quorum.ReadFunc[Key, Data] {
	latestDataFromQuorum := latestDataFromQuorumOfNodesFunc(readNodesDataToChannel, latestData)

	return func(ctx context.Context, key Key) rslt.Of[Data] {
		return fn.Compose(
			rslt.FmapPartial(fn.Bind2(ctx, key, latestDataFromQuorum)), fn.Ctx(ctx, discovery.Nodes),
		)(key)
	}
}

func latestDataFromQuorumOfNodesFunc[Key, Data any](
	readNodesDataToChannel func(context.Context, Key, []quorum.Node[Key, Data]) <-chan Data,
	latestData func([]Data) Data,
) func(context.Context, Key, []quorum.Node[Key, Data]) rslt.Of[Data] {
	quorumOfData := fn.Compose(chn.FirstNFunc[Data], numberOfQuorum)

	return func(ctx context.Context, key Key, nodes []quorum.Node[Key, Data]) rslt.Of[Data] {
		latestDataFromQuorum := fn.Compose3(
			rslt.Fmap(latestData), quorumOfData(len(nodes)), fn.Bind2(ctx, key, readNodesDataToChannel),
		)
		return latestDataFromQuorum(nodes)
	}
}

var numberOfQuorum = fn.Compose(zd.Successor, zd.Half)
