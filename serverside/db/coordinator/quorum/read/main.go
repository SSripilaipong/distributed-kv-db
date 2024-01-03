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
	)
}

func newFunc[Key, Data any](discovery quorum.Discovery[Key, Data], readNodesToChannel func([]quorum.Node[Key, Data]) <-chan Data) quorum.ReadFunc[Key, Data] {
	return func(ctx context.Context, key Key) rslt.Of[Data] {
		nodes := discovery.Nodes(ctx)
		numberOfNodes := len(nodes.Value())
		readQuorumOfData := fn.Compose(
			rslt.FmapPartial(chn.FirstNFunc[Data](numberOfQuorum(numberOfNodes))), rslt.Fmap(readNodesToChannel),
		)
		readQuorumOfData(nodes)
		return rslt.Error[Data](nil)
	}
}

func numberOfQuorum(n int) int {
	return halfOf(n) + 1
}

func halfOf(n int) int {
	return n / 2
}
