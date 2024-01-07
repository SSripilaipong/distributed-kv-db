package quorum

import (
	"context"
	"distributed-kv-db/common/rslt"
)

type DiscoverNodes[Key, Node any] func(ctx context.Context, key Key) rslt.Of[[]Node]

type ReadNode[Key, Data any] interface {
	Read(context.Context, Key) rslt.Of[Data]
}

type ReadFunc[Key, Data any] func(context.Context, Key) rslt.Of[Data]
type WriteFunc[Key, Data any] func(context.Context, Key, Data) error
