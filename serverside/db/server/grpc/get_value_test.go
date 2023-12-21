package grpc

import (
	"distributed-kv-db/api/grpc"
	"distributed-kv-db/common/result"
	"distributed-kv-db/serverside/db/coordinator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_get_value(t *testing.T) {
	t.Run("should call get value with request", func(t *testing.T) {
		var receivedRequest coordinator.GetValueRequest

		runServerAndGetValueWithRequest(
			newWithGetValueFunc(getValueCaptureRequest(&receivedRequest)),
			&grpc.GetValueRequest{Key: "abc"},
		)

		assert.Equal(t, coordinator.GetValueRequest{Key: "abc"}, receivedRequest)
	})

	t.Run("should return response", func(t *testing.T) {
		response, err := runServerAndGetValueWithResponse(
			newWithGetValueFunc(getValueWithResponse(result.Value(coordinator.GetValueResponse{
				Value: "123",
			}))),
		)

		assert.Equal(t, "123", response.Value)
		assert.NoError(t, err)
	})
}
