package grpc

import (
	"distributed-kv-db/api/grpc"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/serverside/db/coordinator/usecase/setvalue"
	"errors"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func Test_set_value(t *testing.T) {
	t.Run("should call set value with request", func(t *testing.T) {
		var receivedRequest setvalue.Request

		runServerAndSetValueWithRequest(
			newWithSetValueFunc(setValueCaptureRequest(&receivedRequest)),
			&grpc.SetValueRequest{
				Key:   "abc",
				Value: "123",
			},
		)

		assert.Equal(t, "abc", receivedRequest.Key)
		assert.Equal(t, "123", receivedRequest.Value)
	})

	t.Run("should return response", func(t *testing.T) {
		response, err := runServerAndSetValueWithResponse(
			newWithSetValueFunc(setValueWithResponse(rslt.Value(setvalue.Response{}))),
		)

		assert.NotNil(t, response) // currently no properties so just check not nil
		assert.NoError(t, err)
	})

	t.Run("should return unknown error", func(t *testing.T) {
		response, err := runServerAndSetValueWithResponse(
			newWithSetValueFunc(setValueWithResponse(rslt.Error[setvalue.Response](errors.New("boom")))),
		)

		assert.Nil(t, response)
		assert.Equal(t, status.Error(codes.Unknown, "boom"), err)
	})
}
