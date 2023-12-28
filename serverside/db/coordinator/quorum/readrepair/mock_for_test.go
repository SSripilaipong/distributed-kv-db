package readrepair

import (
	"context"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/common/typ"
	"distributed-kv-db/serverside/db/coordinator/quorum"
)

func quorumRead[Key, Data any]() quorum.ReadFunc[Key, Data] {
	return func(ctx context.Context, k Key) rslt.Of[Data] {
		return rslt.Value(typ.Zero[Data]())
	}
}

func quorumReadCaptureKey[Key, Data any](key *Key) quorum.ReadFunc[Key, Data] {
	return func(ctx context.Context, k Key) rslt.Of[Data] {
		*key = k
		return rslt.Value(typ.Zero[Data]())
	}
}

func quorumReadCaptureContext[Key, Data any](ctx *context.Context) quorum.ReadFunc[Key, Data] {
	return func(c context.Context, k Key) rslt.Of[Data] {
		*ctx = c
		return rslt.Value(typ.Zero[Data]())
	}
}

func quorumReadWithResult[Key, Data any](result rslt.Of[Data]) quorum.ReadFunc[Key, Data] {
	return func(ctx context.Context, k Key) rslt.Of[Data] {
		return result
	}
}

func quorumWrite[Key, Data any]() quorum.WriteFunc[Key, Data] {
	return func(ctx context.Context, k Key, d Data) error {
		return nil
	}
}

func quorumWriteCaptureKey[Key, Data any](key *Key) quorum.WriteFunc[Key, Data] {
	return func(ctx context.Context, k Key, d Data) error {
		*key = k
		return nil
	}
}

func quorumWriteCaptureData[Key, Data any](data *Data) quorum.WriteFunc[Key, Data] {
	return func(ctx context.Context, k Key, d Data) error {
		*data = d
		return nil
	}
}
