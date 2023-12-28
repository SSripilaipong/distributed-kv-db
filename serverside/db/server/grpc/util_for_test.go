package grpc

import (
	"context"
	"distributed-kv-db/api/grpc"
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/grpcutil"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/serverside/db/coordinator"
)

func setValueCaptureRequest(request *coordinator.SetValueRequest) coordinator.SetValueFunc {
	return func(ctx context.Context, r coordinator.SetValueRequest) rslt.Of[coordinator.SetValueResponse] {
		*request = r
		return rslt.Value(coordinator.SetValueResponse{})
	}
}

func setValueWithResponse(response rslt.Of[coordinator.SetValueResponse]) coordinator.SetValueFunc {
	return func(ctx context.Context, r coordinator.SetValueRequest) rslt.Of[coordinator.SetValueResponse] {
		return response
	}
}

func getValueCaptureRequest(request *coordinator.GetValueRequest) coordinator.GetValueFunc {
	return func(ctx context.Context, r coordinator.GetValueRequest) rslt.Of[coordinator.GetValueResponse] {
		*request = r
		return rslt.Value(coordinator.GetValueResponse{})
	}
}

func getValueWithResponse(response rslt.Of[coordinator.GetValueResponse]) coordinator.GetValueFunc {
	return func(ctx context.Context, r coordinator.GetValueRequest) rslt.Of[coordinator.GetValueResponse] {
		return response
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

func runServerAndGetValueWithResponse(runner Func) (resp *grpc.GetValueResponse, err error) {
	runServerAndExecuteClient(runner, func(client grpc.ServerClient) {
		resp, err = client.GetValue(context.Background(), &grpc.GetValueRequest{})
	})
	return
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
		rslt.Fmap(grpcutil.CloseClient)(conn)
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

var newClient = rslt.Fmap(fn.Compose(grpc.NewServerClient, grpcutil.ClientToInterface))
