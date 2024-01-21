package readlatest

import (
	"context"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/common/typ"
	peerRead "distributed-kv-db/serverside/db/coordinator/quorum/read"
)

func read[Key, Data any](read peerRead.Func[Key, Data]) rslt.Of[Data] {
	return read(context.Background(), typ.Zero[Key]())
}

func readWithContext[Key, Data any](read peerRead.Func[Key, Data], ctx context.Context) rslt.Of[Data] {
	return read(ctx, typ.Zero[Key]())
}

func readWithKey[Key, Data any](read peerRead.Func[Key, Data], key Key) rslt.Of[Data] {
	return read(context.Background(), key)
}
