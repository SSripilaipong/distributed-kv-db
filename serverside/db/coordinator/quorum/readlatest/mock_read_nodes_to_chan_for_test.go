package readlatest

import (
	"context"
	"distributed-kv-db/common/cntx"
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/common/typ"
)

func readQuorumDummy[Key, Data any](_ context.Context, _ Key) rslt.Of[[]Data] {
	return rslt.Value(typ.Zero[[]Data]())
}

func readQuorumCaptureContext[Key, Data any](ctx *context.Context) func(context.Context, Key) rslt.Of[[]Data] {
	return func(c context.Context, _ Key) rslt.Of[[]Data] {
		*ctx = c
		return rslt.Value(typ.Zero[[]Data]())
	}
}

func readQuorumCaptureKey[Key, Data any](key *Key) func(context.Context, Key) rslt.Of[[]Data] {
	return func(_ context.Context, k Key) rslt.Of[[]Data] {
		*key = k
		return rslt.Value(typ.Zero[[]Data]())
	}
}

func readQuorumWithResult[Key, Data any](result rslt.Of[[]Data]) func(context.Context, Key) rslt.Of[[]Data] {
	return cntx.Func(fn.Const[Key](result))
}
