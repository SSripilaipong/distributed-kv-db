package grpc

import (
	"context"
	"distributed-kv-db/common/cntx"
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/serverside/db/coordinator"
)

func getValueCaptureRequest(request *coordinator.GetValueRequest) coordinator.GetValueFunc {
	return func(ctx context.Context, r coordinator.GetValueRequest) rslt.Of[coordinator.GetValueResponse] {
		*request = r
		return rslt.Value(coordinator.GetValueResponse{})
	}
}

func getValueWithResponse(response rslt.Of[coordinator.GetValueResponse]) coordinator.GetValueFunc {
	return cntx.Func(fn.Const[coordinator.GetValueRequest](response))
}
