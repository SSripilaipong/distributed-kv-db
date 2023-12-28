package coordinator

import (
	"context"
	"distributed-kv-db/common/rslt"
)

type GetValueRequest struct {
	Key string
}

type SetValueRequest struct {
	Key   string
	Value string
}

type GetValueResponse struct {
	Value string
}

type SetValueResponse struct{}

type GetValueFunc func(ctx context.Context, request GetValueRequest) rslt.Of[GetValueResponse]

type SetValueFunc func(ctx context.Context, request SetValueRequest) rslt.Of[SetValueResponse]
