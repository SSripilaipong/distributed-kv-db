package read

import (
	"context"
	"distributed-kv-db/common/cntx"
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/serverside/db/coordinator/quorum"
)

func discoverNodesFuncDummy[Key, Data any](_ context.Context, _ Key) rslt.Of[[]quorum.ReadableNode[Key, Data]] {
	return rslt.Value([]quorum.ReadableNode[Key, Data]{})
}

func discoverNodesFuncWithResult[Key, Data any](nodes rslt.Of[[]quorum.ReadableNode[Key, Data]]) quorum.DiscoverNodes[Key, quorum.ReadableNode[Key, Data]] {
	return cntx.Func(fn.Const[Key](nodes))
}

func discoverNodesFuncCaptureContext[Key, Data any](ctx *context.Context) quorum.DiscoverNodes[Key, quorum.ReadableNode[Key, Data]] {
	return func(c context.Context, key Key) rslt.Of[[]quorum.ReadableNode[Key, Data]] {
		*ctx = c
		return rslt.Value([]quorum.ReadableNode[Key, Data]{})
	}
}

func discoverNodesFuncCaptureKey[Key, Data any](key *Key) quorum.DiscoverNodes[Key, quorum.ReadableNode[Key, Data]] {
	return func(c context.Context, k Key) rslt.Of[[]quorum.ReadableNode[Key, Data]] {
		*key = k
		return rslt.Value([]quorum.ReadableNode[Key, Data]{})
	}
}
