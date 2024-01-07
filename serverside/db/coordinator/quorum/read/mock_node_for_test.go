package read

import (
	"context"
	"distributed-kv-db/common/cntx"
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/common/typ"
	"distributed-kv-db/serverside/db/coordinator/quorum"
)

type nodeWithId[Key, Data any] struct {
	id int
}

func newNodeWithId[Key, Data any](id int) nodeWithId[Key, Data] {
	return nodeWithId[Key, Data]{id: id}
}

func (n nodeWithId[Key, Data]) Read(_ context.Context, _ Key) rslt.Of[Data] {
	return rslt.Value(typ.Zero[Data]())
}

var _ quorum.Node[int, int] = nodeWithId[int, int]{}

type nodeMock[Key, Data any] struct {
	read func(context.Context, Key) rslt.Of[Data]
}

func nodeDummy[Key, Data any]() nodeMock[Key, Data] {
	return nodeMock[Key, Data]{read: readFuncDummy[Key, Data]}
}

func (n nodeMock[Key, Data]) Read(ctx context.Context, key Key) rslt.Of[Data] {
	return n.read(ctx, key)
}

func nodeWithReadFunc[Key, Data any](read func(context.Context, Key) rslt.Of[Data]) nodeMock[Key, Data] {
	return nodeMock[Key, Data]{read: read}
}

func readFuncDummy[Key, Data any](context.Context, Key) rslt.Of[Data] {
	return rslt.Value(typ.Zero[Data]())
}

func readFuncCaptureContext[Key, Data any](ctx *context.Context) func(context.Context, Key) rslt.Of[Data] {
	return func(c context.Context, _ Key) rslt.Of[Data] {
		*ctx = c
		return rslt.Value(typ.Zero[Data]())
	}
}

func readFuncCaptureKey[Key, Data any](key *Key) func(context.Context, Key) rslt.Of[Data] {
	return func(_ context.Context, k Key) rslt.Of[Data] {
		*key = k
		return rslt.Value(typ.Zero[Data]())
	}
}

func readFuncWithResult[Key, Data any](result rslt.Of[Data]) func(context.Context, Key) rslt.Of[Data] {
	return cntx.Func(fn.Const[Key](result))
}
