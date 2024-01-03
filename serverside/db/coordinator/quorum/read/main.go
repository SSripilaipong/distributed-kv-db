package read

import (
	"context"
	"distributed-kv-db/common/chn"
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/rslt"
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
	firstQ := fn.Compose(chn.FirstNFunc[Data], numberOfQuorum)

	return func(ctx context.Context, key Key) rslt.Of[Data] {
		nodes := discovery.Nodes(ctx, key)
		numberOfNodes := len(nodes.Value())

		readLatestDataFromQuorum := fn.Compose3(
			rslt.Fmap(latestData), rslt.FmapPartial(firstQ(numberOfNodes)), rslt.Fmap(readNodeDataToChannel),
		)
		return readLatestDataFromQuorum(nodes)
	}
}

func numberOfQuorum(n int) int {
	return halfOf(n) + 1
}

func halfOf(n int) int {
	return n / 2
}
