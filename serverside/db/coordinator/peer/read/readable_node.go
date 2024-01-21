package read

import (
	"context"
	"distributed-kv-db/common/rslt"
)

type ReadableNode[Key, Data any] interface {
	Read(context.Context, Key) rslt.Of[Data]
}
