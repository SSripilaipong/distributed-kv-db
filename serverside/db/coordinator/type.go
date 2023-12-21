package coordinator

import (
	"context"
	"distributed-kv-db/common/result"
)

type GetValueRequest struct {
	Key string
}

type GetValueResponse struct {
	Value string
}

type SetValueRequest struct {
	Key   string
	Value string
}

type SetValueResponse struct{}

type GetValueFunc func(ctx context.Context, request GetValueRequest) result.Of[GetValueResponse]

type SetValueFunc func(ctx context.Context, request SetValueRequest) result.Of[SetValueResponse]
