package getvalue

import (
	"context"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/serverside/db/coordinator"
)

func getValueWithRequest(getValue coordinator.GetValueFunc, request coordinator.GetValueRequest) {
	getValue(context.Background(), request)
}

func getValueWithContext(getValue coordinator.GetValueFunc, ctx context.Context) {
	getValue(ctx, coordinator.GetValueRequest{})
}

func getValueWithResponse(getValue coordinator.GetValueFunc) rslt.Of[coordinator.GetValueResponse] {
	return getValue(context.Background(), coordinator.GetValueRequest{})
}
