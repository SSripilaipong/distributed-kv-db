package setvalue

import (
	"context"
	"distributed-kv-db/common/rslt"
)

type Func func(ctx context.Context, request Request) rslt.Of[Response]
