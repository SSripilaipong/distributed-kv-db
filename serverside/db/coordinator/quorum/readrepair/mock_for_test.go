package readrepair

import (
	"context"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/common/typ"
	"distributed-kv-db/serverside/db/coordinator/quorum"
)

func readQuorumCaptureKey[Key, Data any](key *Key) quorum.ReadFunc[Key, Data] {
	return func(ctx context.Context, k Key) rslt.Of[Data] {
		*key = k
		return rslt.Value(typ.Zero[Data]())
	}
}
