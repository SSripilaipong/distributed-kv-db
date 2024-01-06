package read

import (
	"context"
	"distributed-kv-db/common/chn"
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/common/slc"
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
	quorumDataOfN := fn.Compose(chn.FirstNFunc[Data], numberOfQuorum)

	return func(ctx context.Context, key Key) rslt.Of[Data] {
		nodes := discovery.Nodes(ctx, key)
		quorumData := quorumDataOfN(countResult(nodes))

		latestDataFromQuorum := fn.Compose3(
			rslt.Fmap(latestData), rslt.FmapPartial(quorumData), rslt.Fmap(readNodeDataToChannel),
		)
		return latestDataFromQuorum(nodes)
	}
}

var numberOfQuorum = fn.Compose(zd.Successor, zd.Half)

func countResult[T any](xs rslt.Of[[]T]) int {
	return rslt.Resolve(slc.Len[T], fn.Const[error](0), xs)
}
