package grpc

import (
	"distributed-kv-db/api/grpc"
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
