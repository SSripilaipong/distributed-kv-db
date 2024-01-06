package grpc

import (
	"context"
	"distributed-kv-db/api/grpc"
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/grpcutil"
	"distributed-kv-db/common/rslt"
)

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

var newClient = rslt.Fmap(fn.Compose(grpc.NewServerClient, grpcutil.ClientToInterface))
