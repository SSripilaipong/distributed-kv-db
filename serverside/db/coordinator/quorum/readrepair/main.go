package readrepair

import (
	"context"
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/serverside/db/coordinator/peer/discovery"
	peerRead "distributed-kv-db/serverside/db/coordinator/peer/read"
	"distributed-kv-db/serverside/db/coordinator/quorum/read"
	"distributed-kv-db/serverside/db/coordinator/quorum/readlatest"
	"distributed-kv-db/serverside/db/coordinator/quorum/write"
	"distributed-kv-db/serverside/db/data/temporal"
)

func New[Key any, Data temporal.Hashable](nReplicas uint, discoverNodes discovery.Func[Key, peerRead.ReadableNode[Key, Data]]) read.Func[Key, Data] {
	return newFunc[Key, Data](
		readlatest.New[Key, Data](read.NodesToDataSlice(nReplicas, discoverNodes)),
		write.New(discoverNodes),
	)
}

func newFunc[Key, Data any](quorumRead read.Func[Key, Data], quorumWrite write.Func[Key, Data]) read.Func[Key, Data] {
	return func(ctx context.Context, key Key) rslt.Of[Data] {
		readResult := quorumRead(ctx, key)
		quorumWriteIfNotError := fmapResultToError(fn.WithArg2(ctx, key, quorumWrite))
		return rslt.ResultOrError(readResult, quorumWriteIfNotError(readResult))
	}
}

func fmapResultToError[A any](f func(A) error) func(rslt.Of[A]) error {
	return fn.Compose(rslt.ErrorOf, rslt.Fmap(f))
}
