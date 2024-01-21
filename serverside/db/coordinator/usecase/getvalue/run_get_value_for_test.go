package getvalue

import (
	"context"
	"distributed-kv-db/common/rslt"
)

func getValueWithRequest(getValue Func, request Request) {
	getValue(context.Background(), request)
}

func getValueWithContext(getValue Func, ctx context.Context) {
	getValue(ctx, Request{})
}

func getValueWithResponse(getValue Func) rslt.Of[Response] {
	return getValue(context.Background(), Request{})
}
