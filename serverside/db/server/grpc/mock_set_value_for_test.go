package grpc

import (
	"context"
	"distributed-kv-db/common/cntx"
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/serverside/db/coordinator/usecase/setvalue"
)

func setValueCaptureRequest(request *setvalue.Request) setvalue.Func {
	return func(ctx context.Context, r setvalue.Request) rslt.Of[setvalue.Response] {
		*request = r
		return rslt.Value(setvalue.Response{})
	}
}

func setValueWithResponse(response rslt.Of[setvalue.Response]) setvalue.Func {
	return cntx.Func(fn.Const[setvalue.Request](response))
}
