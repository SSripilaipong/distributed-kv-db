package setvalue

import (
	"context"
	"distributed-kv-db/common/rslt"
)

func New() Func {
	return func(ctx context.Context, request Request) rslt.Of[Response] {
		return rslt.Value(Response{})
	}
}
