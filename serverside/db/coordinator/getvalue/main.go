package getvalue

import (
	"context"
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/result"
	"distributed-kv-db/serverside/db/coordinator"
)

var New = fn.Compose(newFunc, newReadRepairFunc)

func newFunc(readRepair readRepairFunc) coordinator.GetValueFunc {
	valueToResponse := result.Fmap(valueToResponse)

	return func(ctx context.Context, request coordinator.GetValueRequest) result.Of[coordinator.GetValueResponse] {
		readRepair := fn.Ctx(ctx, readRepair)

		return valueToResponse(readRepair(requestToKey(request)))
	}
}
