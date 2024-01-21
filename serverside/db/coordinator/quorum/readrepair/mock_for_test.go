package readrepair

import (
	"context"
	"distributed-kv-db/common/cntx"
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/common/typ"
	"distributed-kv-db/serverside/db/coordinator/quorum/read"
	"distributed-kv-db/serverside/db/coordinator/quorum/write"
)

func quorumRead[Key, Data any]() read.Func[Key, Data] {
	return cntx.Func(fn.Const[Key](rslt.Value(typ.Zero[Data]())))
}

func quorumReadCaptureKey[Key, Data any](key *Key) read.Func[Key, Data] {
	return func(ctx context.Context, k Key) rslt.Of[Data] {
		*key = k
		return rslt.Value(typ.Zero[Data]())
	}
}

func quorumReadCaptureContext[Key, Data any](ctx *context.Context) read.Func[Key, Data] {
	return func(c context.Context, k Key) rslt.Of[Data] {
		*ctx = c
		return rslt.Value(typ.Zero[Data]())
	}
}

func quorumReadWithResult[Key, Data any](result rslt.Of[Data]) read.Func[Key, Data] {
	return cntx.Func(fn.Const[Key](result))
}

func quorumWrite[Key, Data any]() write.Func[Key, Data] {
	return cntx.Func2(fn.Const2[Key, Data, error](nil))
}

func quorumWriteCaptureKey[Key, Data any](key *Key) write.Func[Key, Data] {
	return func(ctx context.Context, k Key, d Data) error {
		*key = k
		return nil
	}
}

func quorumWriteCaptureData[Key, Data any](data *Data) write.Func[Key, Data] {
	return func(ctx context.Context, k Key, d Data) error {
		*data = d
		return nil
	}
}

func quorumWriteCaptureContext[Key, Data any](ctx *context.Context) write.Func[Key, Data] {
	return func(c context.Context, k Key, d Data) error {
		*ctx = c
		return nil
	}
}

func quorumWriteCaptureIsCalled[Key, Data any](isCalled *bool) write.Func[Key, Data] {
	return func(c context.Context, k Key, d Data) error {
		*isCalled = true
		return nil
	}
}

func quorumWriteWithError[Key, Data any](err error) write.Func[Key, Data] {
	return cntx.Func2(fn.Const2[Key, Data](err))
}
