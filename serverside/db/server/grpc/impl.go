package grpc

import (
	"context"
	"distributed-kv-db/api/grpc"
	usecaseError "distributed-kv-db/serverside/db/coordinator/usecase/error"
	"distributed-kv-db/serverside/db/coordinator/usecase/getvalue"
	"distributed-kv-db/serverside/db/coordinator/usecase/setvalue"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type grpcImpl struct {
	grpc.UnimplementedServerServer
	getValue getvalue.Func
	setValue setvalue.Func
}

func (t grpcImpl) GetValue(ctx context.Context, req *grpc.GetValueRequest) (*grpc.GetValueResponse, error) {
	resp := t.getValue(ctx, getvalue.Request{Key: req.Key})
	if resp.IsOk() {
		return &grpc.GetValueResponse{Value: resp.Value().Value}, nil
	}
	if errors.As(resp.Error(), &usecaseError.KeyNotFound{}) {
		return nil, status.Error(codes.NotFound, resp.Error().Error())
	}
	return nil, status.Error(codes.Unknown, resp.Error().Error())
}

func (t grpcImpl) SetValue(ctx context.Context, req *grpc.SetValueRequest) (*grpc.SetValueResponse, error) {
	resp := t.setValue(ctx, setvalue.Request{
		Key:   req.Key,
		Value: req.Value,
	})
	if resp.IsOk() {
		return &grpc.SetValueResponse{}, nil
	}
	return nil, status.Error(codes.Unknown, resp.Error().Error())
}
