package grpc

import (
	"context"
	"distributed-kv-db/api/grpc"
	"distributed-kv-db/serverside/db/coordinator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type grpcImpl struct {
	grpc.UnimplementedServerServer
	getValue coordinator.GetValueFunc
	setValue coordinator.SetValueFunc
}

func (t grpcImpl) GetValue(ctx context.Context, req *grpc.GetValueRequest) (*grpc.GetValueResponse, error) {
	t.getValue(ctx, coordinator.GetValueRequest{Key: req.Key})
	return nil, nil
}

func (t grpcImpl) SetValue(ctx context.Context, req *grpc.SetValueRequest) (*grpc.SetValueResponse, error) {
	resp := t.setValue(ctx, coordinator.SetValueRequest{
		Key:   req.Key,
		Value: req.Value,
	})
	if resp.IsOk() {
		return &grpc.SetValueResponse{}, nil
	}
	return nil, status.Error(codes.Unknown, resp.Error().Error())
}
