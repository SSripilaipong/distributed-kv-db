package grpc

import (
	"context"
	"distributed-kv-db/api/grpc"
	"distributed-kv-db/serverside/db/coordinator"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type grpcImpl struct {
	grpc.UnimplementedServerServer
	getValue coordinator.GetValueFunc
	setValue coordinator.SetValueFunc
}

func (t grpcImpl) GetValue(ctx context.Context, req *grpc.GetValueRequest) (*grpc.GetValueResponse, error) {
	resp := t.getValue(ctx, getValueRequest{key: req.Key})
	if resp.IsOk() {
		return &grpc.GetValueResponse{Value: resp.Value().Value}, nil
	}
	if errors.As(resp.Error(), &coordinator.KeyNotFoundError{}) {
		return nil, status.Error(codes.NotFound, resp.Error().Error())
	}
	return nil, status.Error(codes.Unknown, resp.Error().Error())
}

func (t grpcImpl) SetValue(ctx context.Context, req *grpc.SetValueRequest) (*grpc.SetValueResponse, error) {
	resp := t.setValue(ctx, setValueRequest{
		key:   req.Key,
		value: req.Value,
	})
	if resp.IsOk() {
		return &grpc.SetValueResponse{}, nil
	}
	return nil, status.Error(codes.Unknown, resp.Error().Error())
}
