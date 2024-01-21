package read

import (
	"context"
	"distributed-kv-db/common/rslt"
)

func discoverNodesDummy[Key, Node any](_ context.Context, _ Key) rslt.Of[[]Node] {
	return rslt.Value([]Node{})
}

func discoverNodesCaptureKey[Key, Node any](key *Key) func(ctx context.Context, key Key) rslt.Of[[]Node] {
	return func(c context.Context, k Key) rslt.Of[[]Node] {
		*key = k
		return discoverNodesDummy[Key, Node](c, k)
	}
}

func discoverNodesCaptureContext[Key, Node any](ctx *context.Context) func(ctx context.Context, key Key) rslt.Of[[]Node] {
	return func(c context.Context, k Key) rslt.Of[[]Node] {
		*ctx = c
		return discoverNodesDummy[Key, Node](c, k)
	}
}

func discoverNodesWithResult[Key, Node any](result rslt.Of[[]Node]) func(ctx context.Context, key Key) rslt.Of[[]Node] {
	return func(c context.Context, k Key) rslt.Of[[]Node] {
		return result
	}
}
