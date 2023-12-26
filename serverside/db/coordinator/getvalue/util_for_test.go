package getvalue

import (
	"context"
	"distributed-kv-db/common/result"
)

var newFuncWithReadRepairFunc = newFunc

func readRepairCaptureQuery(query *string) readRepairFunc {
	return func(ctx context.Context, q string) result.Of[string] {
		*query = q
		return result.Value("")
	}
}

func readRepairCaptureContext(ctx *context.Context) readRepairFunc {
	return func(c context.Context, q string) result.Of[string] {
		*ctx = c
		return result.Value("")
	}
}

func readRepairWithResponse(resp result.Of[string]) readRepairFunc {
	return func(c context.Context, q string) result.Of[string] {
		return resp
	}
}
