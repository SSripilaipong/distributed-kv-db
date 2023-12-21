package grpc

import (
	"context"
	"distributed-kv-db/api/grpc"
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/grpcutil"
	"distributed-kv-db/common/result"
	"distributed-kv-db/serverside/db/coordinator"
)

func setValueCaptureRequest(request *coordinator.SetValueRequest) coordinator.SetValueFunc {
	return func(ctx context.Context, r coordinator.SetValueRequest) result.Of[coordinator.SetValueResponse] {
		*request = r
		return result.Value(coordinator.SetValueResponse{})
	}
}

func setValueWithReturn(response result.Of[coordinator.SetValueResponse]) coordinator.SetValueFunc {
	return func(ctx context.Context, r coordinator.SetValueRequest) result.Of[coordinator.SetValueResponse] {
		return response
	}
}

func getValueCaptureRequest(request *coordinator.GetValueRequest) coordinator.GetValueFunc {
	return func(ctx context.Context, r coordinator.GetValueRequest) result.Of[coordinator.GetValueResponse] {
		*request = r
		return result.Value(coordinator.GetValueResponse{})
	}
}

func runServerAndSetValueWithResponse(runner Func) (resp *grpc.SetValueResponse, err error) {
	runServerAndExecuteClient(runner, func(client grpc.ServerClient) {
		resp, err = client.SetValue(context.Background(), &grpc.SetValueRequest{})
	})
	return
}

func runServerAndSetValueWithRequest(runner Func, request *grpc.SetValueRequest) {
	runServerAndExecuteClient(runner, func(client grpc.ServerClient) {
		_, _ = client.SetValue(context.Background(), request)
	})
}

func runServerAndGetValueWithRequest(runner Func, request *grpc.GetValueRequest) {
	runServerAndExecuteClient(runner, func(client grpc.ServerClient) {
		_, _ = client.GetValue(context.Background(), request)
	})
}

func runServerAndExecuteClient(grpcRunner Func, clientExecute func(client grpc.ServerClient)) {
	server := grpcRunner(0).Value()

	conn := grpcutil.Connect(grpcutil.LocalAddress(server.Port()))
	client := newClient(conn).Value()

	go func() {
		clientExecute(client)
		result.Fmap(grpcutil.CloseClient)(conn)
		server.ForceStop()
	}()
	<-server.Done()
}

func newWithSetValueFunc(setValue coordinator.SetValueFunc) Func {
	return New(nil, setValue)
}

func newWithGetValueFunc(getValue coordinator.GetValueFunc) Func {
	return New(getValue, nil)
}

var newClient = result.Fmap(fn.Compose(grpc.NewServerClient, grpcutil.ClientToInterface))
