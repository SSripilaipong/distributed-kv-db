package getvalue

import (
	"context"
	"distributed-kv-db/common/result"
	"distributed-kv-db/serverside/db/coordinator"
)

func New(discovery quorumDiscovery) coordinator.GetValueFunc {
	return newFunc(newReadRepairFunc(discovery))
}

func newFunc(readRepair readRepairFunc) coordinator.GetValueFunc {
	return func(ctx context.Context, request coordinator.GetValueRequest) result.Of[coordinator.GetValueResponse] {
		readRepair(ctx, request.Key)
		return result.Value(coordinator.GetValueResponse{Value: "tmp"})
	}
}
