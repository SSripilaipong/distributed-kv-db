package read

import (
	"context"
	"distributed-kv-db/common/chn"
	"distributed-kv-db/common/cntx"
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/typ"
	"distributed-kv-db/serverside/db/coordinator/quorum"
)

func readNodesDataToChannelDummy[Key, Data any](_ context.Context, _ []quorum.Node[Key, Data]) <-chan Data {
	return chn.Repeat(typ.Zero[Data]())
}

func readNodesDataToChannelCaptureContext[Key, Data any](ctx *context.Context) func(context.Context, []quorum.Node[Key, Data]) <-chan Data {
	return func(c context.Context, _ []quorum.Node[Key, Data]) <-chan Data {
		*ctx = c
		return chn.Repeat(typ.Zero[Data]())
	}
}

func readNodesDataToChannelCaptureNodes[Key, Data any](nodes *[]quorum.Node[Key, Data]) func(context.Context, []quorum.Node[Key, Data]) <-chan Data {
	return func(_ context.Context, n []quorum.Node[Key, Data]) <-chan Data {
		*nodes = n
		return chn.Repeat(typ.Zero[Data]())
	}
}

func readNodesDataToChannelWithResult[Key, Data any](ch <-chan Data) func(context.Context, []quorum.Node[Key, Data]) <-chan Data {
	return cntx.Func(fn.Const[[]quorum.Node[Key, Data]](ch))
}
