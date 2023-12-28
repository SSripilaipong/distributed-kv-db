package quorum

import (
	"context"
	"distributed-kv-db/common/result"
)

type Discovery[Key, Data any] interface {
}

type ReadFunc[Key, Data any] func(context.Context, Key) result.Of[Data]
type WriteFunc[Key, Data any] func(context.Context, Key, Data) error
