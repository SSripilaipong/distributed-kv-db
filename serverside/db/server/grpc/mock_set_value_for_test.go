package grpc

import (
	"context"
	"distributed-kv-db/common/cntx"
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/serverside/db/coordinator"
)

func setValueCaptureRequest(request *coordinator.SetValueRequest) coordinator.SetValueFunc {
	return func(ctx context.Context, r coordinator.SetValueRequest) rslt.Of[coordinator.SetValueResponse] {
		*request = r
		return rslt.Value(coordinator.SetValueResponse{})
	}
}

func setValueWithResponse(response rslt.Of[coordinator.SetValueResponse]) coordinator.SetValueFunc {
	return cntx.Func(fn.Const[coordinator.SetValueRequest](response))
}
