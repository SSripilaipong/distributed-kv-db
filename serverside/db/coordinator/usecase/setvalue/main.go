package setvalue

import (
	"context"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/serverside/db/coordinator"
)

func New() coordinator.SetValueFunc {
	return func(ctx context.Context, request coordinator.SetValueRequest) rslt.Of[coordinator.SetValueResponse] {
		return rslt.Value(coordinator.SetValueResponse{})
	}
}
