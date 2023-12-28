package quorum

import (
	"context"
	"distributed-kv-db/common/rslt"
)

type Discovery[Key, Data any] interface {
}

type ReadFunc[Key, Data any] func(context.Context, Key) rslt.Of[Data]
type WriteFunc[Key, Data any] func(context.Context, Key, Data) error
