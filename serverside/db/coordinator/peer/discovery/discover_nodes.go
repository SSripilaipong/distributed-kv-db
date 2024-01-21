package discovery

import (
	"context"
	"distributed-kv-db/common/rslt"
)

type DiscoverNodes[Key, Node any] func(ctx context.Context, key Key) rslt.Of[[]Node]
