package readlatest

import (
	"context"
	"distributed-kv-db/common/cntx"
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/common/typ"
)

func readQuorumOfNodesDataDummy[Key, Data, Node any](_ context.Context, _ Key, _ []Node) rslt.Of[[]Data] {
	return rslt.Value(typ.Zero[[]Data]())
}

func readQuorumOfNodesDataCaptureContext[Key, Data, Node any](ctx *context.Context) func(context.Context, Key, []Node) rslt.Of[[]Data] {
	return func(c context.Context, _ Key, _ []Node) rslt.Of[[]Data] {
		*ctx = c
		return rslt.Value(typ.Zero[[]Data]())
	}
}

func readQuorumOfNodesDataCaptureKey[Key, Data, Node any](key *Key) func(context.Context, Key, []Node) rslt.Of[[]Data] {
	return func(_ context.Context, k Key, _ []Node) rslt.Of[[]Data] {
		*key = k
		return rslt.Value(typ.Zero[[]Data]())
	}
}

func readQuorumOfNodesDataChannelCaptureNodes[Key, Data, Node any](nodes *[]Node) func(context.Context, Key, []Node) rslt.Of[[]Data] {
	return func(_ context.Context, _ Key, n []Node) rslt.Of[[]Data] {
		*nodes = n
		return rslt.Value(typ.Zero[[]Data]())
	}
}

func readQuorumOfNodesDataWithResult[Key, Data, Node any](result rslt.Of[[]Data]) func(context.Context, Key, []Node) rslt.Of[[]Data] {
	return cntx.Func2(fn.Const2[Key, []Node](result))
}
