package grpc

import (
	"distributed-kv-db/api/grpc"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/serverside/db/coordinator"
	"errors"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func Test_get_value(t *testing.T) {
	t.Run("should call get value with request", func(t *testing.T) {
		var receivedRequest coordinator.GetValueRequest

		runServerAndGetValueWithRequest(
			newWithGetValueFunc(getValueCaptureRequest(&receivedRequest)),
			&grpc.GetValueRequest{Key: "abc"},
		)

		assert.Equal(t, "abc", receivedRequest.Key)
	})

	t.Run("should return response", func(t *testing.T) {
		response, err := runServerAndGetValueWithResponse(
			newWithGetValueFunc(getValueWithResponse(rslt.Value(coordinator.GetValueResponse{
				Value: "123",
			}))),
		)

		assert.Equal(t, "123", response.GetValue())
		assert.NoError(t, err)
	})

	t.Run("should return not found error", func(t *testing.T) {
		response, err := runServerAndGetValueWithResponse(
			newWithGetValueFunc(getValueWithResponse(
				rslt.Error[coordinator.GetValueResponse](coordinator.NewKeyNotFoundError("aaa")),
			)),
		)

		assert.Nil(t, response)
		assert.Equal(t, status.Error(codes.NotFound, "key \"aaa\" not found"), err)
	})

	t.Run("should return unknown error", func(t *testing.T) {
		response, err := runServerAndGetValueWithResponse(
			newWithGetValueFunc(getValueWithResponse(
				rslt.Error[coordinator.GetValueResponse](errors.New("boom")),
			)),
		)

		assert.Nil(t, response)
		assert.Equal(t, status.Error(codes.Unknown, "boom"), err)
	})
}
