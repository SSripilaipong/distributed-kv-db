package read

import (
	"context"
	"distributed-kv-db/common/cntx"
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/serverside/db/coordinator/quorum"
)

func discoverNodesFuncDummy[Key, Node any](_ context.Context, _ Key) rslt.Of[[]Node] {
	return rslt.Value([]Node{})
}

func discoverNodesFuncWithResult[Key, Node any](nodes rslt.Of[[]Node]) quorum.DiscoverNodes[Key, Node] {
	return cntx.Func(fn.Const[Key](nodes))
}

func discoverNodesFuncCaptureContext[Key, Node any](ctx *context.Context) quorum.DiscoverNodes[Key, Node] {
	return func(c context.Context, key Key) rslt.Of[[]Node] {
		*ctx = c
		return rslt.Value([]Node{})
	}
}

func discoverNodesFuncCaptureKey[Key, Node any](key *Key) quorum.DiscoverNodes[Key, Node] {
	return func(c context.Context, k Key) rslt.Of[[]Node] {
		*key = k
		return rslt.Value([]Node{})
	}
}
