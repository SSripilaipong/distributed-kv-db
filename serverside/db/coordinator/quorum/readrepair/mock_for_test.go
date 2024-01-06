package readrepair

import (
	"context"
	"distributed-kv-db/common/cntx"
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/common/typ"
	"distributed-kv-db/serverside/db/coordinator/quorum"
)

func quorumRead[Key, Data any]() quorum.ReadFunc[Key, Data] {
	return cntx.Func(fn.Const[Key](rslt.Value(typ.Zero[Data]())))
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
	return cntx.Func(fn.Const[Key](result))
}

func quorumWrite[Key, Data any]() quorum.WriteFunc[Key, Data] {
	return cntx.Func2(fn.Const2[Key, Data, error](nil))
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

func quorumWriteCaptureContext[Key, Data any](ctx *context.Context) quorum.WriteFunc[Key, Data] {
	return func(c context.Context, k Key, d Data) error {
		*ctx = c
		return nil
	}
}

func quorumWriteCaptureIsCalled[Key, Data any](isCalled *bool) quorum.WriteFunc[Key, Data] {
	return func(c context.Context, k Key, d Data) error {
		*isCalled = true
		return nil
	}
}

func quorumWriteWithError[Key, Data any](err error) quorum.WriteFunc[Key, Data] {
	return cntx.Func2(fn.Const2[Key, Data](err))
}
