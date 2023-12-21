package coordinator

import (
	"context"
	"distributed-kv-db/common/result"
)

func NewGetValue() GetValueFunc {
	return func(ctx context.Context, request GetValueRequest) result.Of[GetValueResponse] {
		return result.Value(GetValueResponse{Value: "tmp"})
	}
}

func NewSetValue() SetValueFunc {
	return func(ctx context.Context, request SetValueRequest) result.Of[SetValueResponse] {
		return result.Value(SetValueResponse{})
	}
}
