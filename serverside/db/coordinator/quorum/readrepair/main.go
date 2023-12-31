package readrepair

import (
	"context"
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/serverside/db/coordinator/quorum"
	"distributed-kv-db/serverside/db/coordinator/quorum/read"
	"distributed-kv-db/serverside/db/coordinator/quorum/write"
)

type orderableData interface {
	quorum.Orderable
	quorum.Hashable
}

func New[Key any, Data orderableData](discoverNodes quorum.DiscoverNodes[Key, quorum.ReadNode[Key, Data]]) quorum.ReadFunc[Key, Data] {
	return newFunc[Key, Data](
		read.New[Key, Data](discoverNodes),
		write.New(discoverNodes),
	)
}

func newFunc[Key, Data any](quorumRead quorum.ReadFunc[Key, Data], quorumWrite quorum.WriteFunc[Key, Data]) quorum.ReadFunc[Key, Data] {
	return func(ctx context.Context, key Key) rslt.Of[Data] {
		readResult := quorumRead(ctx, key)
		quorumWriteIfNotError := fmapResultToError(fn.Bind2(ctx, key, quorumWrite))
		return rslt.ResultOrError(readResult, quorumWriteIfNotError(readResult))
	}
}

func fmapResultToError[A any](f func(A) error) func(rslt.Of[A]) error {
	return fn.Compose(rslt.OfError, rslt.Fmap(f))
}
