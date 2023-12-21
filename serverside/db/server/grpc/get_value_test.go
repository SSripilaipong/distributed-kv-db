package grpc

import (
	"distributed-kv-db/api/grpc"
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
}
