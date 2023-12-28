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

func readQuorumWithResult[Key, Data any](result rslt.Of[Data]) quorum.ReadFunc[Key, Data] {
	return func(ctx context.Context, k Key) rslt.Of[Data] {
		return result
	}
}

func quorumWrite[Key, Data any]() quorum.WriteFunc[Key, Data] {
	return func(ctx context.Context, k Key, d Data) error {
		return nil
	}
}

func writeQuorumCaptureKeyAndData[Key, Data any](key *Key, data *Data) quorum.WriteFunc[Key, Data] {
	return func(ctx context.Context, k Key, d Data) error {
		*key, *data = k, d
		return nil
	}
}
