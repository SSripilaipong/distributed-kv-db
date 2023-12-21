package grpc

import (
	"context"
	"distributed-kv-db/api/grpc"
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/grpcutil"
	"distributed-kv-db/common/result"
	"distributed-kv-db/serverside/db/coordinator"
	"errors"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func Test_set_value(t *testing.T) {
	t.Run("should call set value with request", func(t *testing.T) {
		var receivedRequest coordinator.SetValueRequest

		runServerAndSetValueWithRequest(
			newWithSetValueFunc(setValueCaptureRequest(&receivedRequest)),
			&grpc.SetValueRequest{
				Key:   "abc",
				Value: "123",
			},
		)

		assert.Equal(t, coordinator.SetValueRequest{
			Key:   "abc",
			Value: "123",
		}, receivedRequest)
	})

	t.Run("should return response", func(t *testing.T) {
		response, err := runServerAndSetValueWithResponse(
			newWithSetValueFunc(setValueWithReturn(result.Value(coordinator.SetValueResponse{}))),
		)

		assert.NotNil(t, response) // currently no properties so just check not nil
		assert.NoError(t, err)
	})

	t.Run("should return unknown error", func(t *testing.T) {
		response, err := runServerAndSetValueWithResponse(
			newWithSetValueFunc(setValueWithReturn(result.Error[coordinator.SetValueResponse](errors.New("boom")))),
		)

		assert.Nil(t, response)
		assert.Equal(t, status.Error(codes.Unknown, "boom"), err)
	})
}

func setValueCaptureRequest(request *coordinator.SetValueRequest) coordinator.SetValueFunc {
	return func(r coordinator.SetValueRequest) result.Of[coordinator.SetValueResponse] {
		*request = r
		return result.Value(coordinator.SetValueResponse{})
	}
}

func setValueWithReturn(response result.Of[coordinator.SetValueResponse]) coordinator.SetValueFunc {
	return func(r coordinator.SetValueRequest) result.Of[coordinator.SetValueResponse] {
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

var newClient = result.Fmap(fn.Compose(grpc.NewServerClient, grpcutil.ClientToInterface))
