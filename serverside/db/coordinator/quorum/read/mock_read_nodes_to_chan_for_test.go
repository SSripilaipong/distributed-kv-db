package read

import (
	"context"
	"distributed-kv-db/common/chn"
	"distributed-kv-db/common/cntx"
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/typ"
)

func readNodesDataToChannelDummy[Key, Data, Node any](_ context.Context, _ Key, _ []Node) <-chan Data {
	return chn.Repeat(typ.Zero[Data]())
}

func readNodesDataToChannelCaptureContext[Key, Data, Node any](ctx *context.Context) func(context.Context, Key, []Node) <-chan Data {
	return func(c context.Context, _ Key, _ []Node) <-chan Data {
		*ctx = c
		return chn.Repeat(typ.Zero[Data]())
	}
}

func readNodesDataToChannelCaptureKey[Key, Data, Node any](key *Key) func(context.Context, Key, []Node) <-chan Data {
	return func(_ context.Context, k Key, _ []Node) <-chan Data {
		*key = k
		return chn.Repeat(typ.Zero[Data]())
	}
}

func readNodesDataToChannelCaptureNodes[Key, Data, Node any](nodes *[]Node) func(context.Context, Key, []Node) <-chan Data {
	return func(_ context.Context, _ Key, n []Node) <-chan Data {
		*nodes = n
		return chn.Repeat(typ.Zero[Data]())
	}
}

func readNodesDataToChannelWithResult[Key, Data, Node any](ch <-chan Data) func(context.Context, Key, []Node) <-chan Data {
	return cntx.Func2(fn.Const2[Key, []Node](ch))
}
