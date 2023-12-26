package setvalue

import (
	"context"
	"distributed-kv-db/common/result"
	"distributed-kv-db/serverside/db/coordinator"
)

func New() coordinator.SetValueFunc {
	return func(ctx context.Context, request coordinator.SetValueRequest) result.Of[coordinator.SetValueResponse] {
		return result.Value(coordinator.SetValueResponse{})
	}
}
