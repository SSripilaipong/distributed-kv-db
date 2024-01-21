package grpc

import (
	"context"
	"distributed-kv-db/common/cntx"
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/serverside/db/coordinator/usecase/getvalue"
)

func getValueCaptureRequest(request *getvalue.Request) getvalue.Func {
	return func(ctx context.Context, r getvalue.Request) rslt.Of[getvalue.Response] {
		*request = r
		return rslt.Value(getvalue.Response{})
	}
}

func getValueWithResponse(response rslt.Of[getvalue.Response]) getvalue.Func {
	return cntx.Func(fn.Const[getvalue.Request](response))
}
