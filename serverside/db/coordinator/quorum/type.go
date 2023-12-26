package quorum

import (
	"context"
	"distributed-kv-db/common/result"
)

type Discovery[Query, Data any] interface {
}

type ReadRepairFunc[Query, Data any] func(context.Context, Query) result.Of[Data]
