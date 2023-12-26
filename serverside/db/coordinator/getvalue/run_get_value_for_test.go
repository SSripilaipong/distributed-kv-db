package getvalue

import (
	"context"
	"distributed-kv-db/serverside/db/coordinator"
)

func getValueWithRequest(getValue coordinator.GetValueFunc, request coordinator.GetValueRequest) {
	getValue(context.Background(), request)
}
