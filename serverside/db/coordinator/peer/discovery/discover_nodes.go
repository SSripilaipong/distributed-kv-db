package discovery

import (
	"context"
	"distributed-kv-db/common/rslt"
)

type Func[Key, Node any] func(ctx context.Context, key Key) rslt.Of[[]Node]
