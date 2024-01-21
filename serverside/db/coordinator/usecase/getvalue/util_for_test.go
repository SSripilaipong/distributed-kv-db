package getvalue

import (
	"context"
	"distributed-kv-db/common/cntx"
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/rslt"
)

var newWithReadRepairFunc = New

func readRepairCaptureQuery(query *string) readRepairFunc {
	return func(ctx context.Context, q string) rslt.Of[string] {
		*query = q
		return rslt.Value("")
	}
}

func readRepairCaptureContext(ctx *context.Context) readRepairFunc {
	return func(c context.Context, q string) rslt.Of[string] {
		*ctx = c
		return rslt.Value("")
	}
}

func readRepairWithResponse(resp rslt.Of[string]) readRepairFunc {
	return cntx.Func(fn.Const[string](resp))
}
