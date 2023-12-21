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
	"testing"
)

func Test_set_value(t *testing.T) {
	t.Run("should call set value with request", func(t *testing.T) {
		var receivedRequest coordinator.SetValueRequest
		grpcRunner := New(nil, setValueCaptureRequest(&receivedRequest))

		_, _ = runServerAndSetValue(grpcRunner, &grpc.SetValueRequest{
			Key:   "abc",
			Value: "123",
		})

		assert.Equal(t, coordinator.SetValueRequest{
			Key:   "abc",
			Value: "123",
		}, receivedRequest)
	})
}

func setValueCaptureRequest(request *coordinator.SetValueRequest) coordinator.SetValueFunc {
	return func(r coordinator.SetValueRequest) result.Of[coordinator.SetValueResponse] {
		*request = r
		return result.Error[coordinator.SetValueResponse](errors.New("stop server"))
	}
}

func runServerAndSetValue(runner Func, request *grpc.SetValueRequest) (resp *grpc.SetValueResponse, err error) {
	runServerAndExecuteClient(runner, func(client grpc.ServerClient) {
		resp, err = client.SetValue(context.Background(), request)
	})
	return
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

var newClient = result.Fmap(fn.Compose(grpc.NewServerClient, grpcutil.ClientToInterface))
