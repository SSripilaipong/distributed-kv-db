package getvalue

import (
	"context"
	"distributed-kv-db/common/cntx"
	"distributed-kv-db/common/result"
	"distributed-kv-db/serverside/db/coordinator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_newFunc(t *testing.T) {
	t.Run("should call read repair with key as query", func(tt *testing.T) {
		var query string

		getValueWithRequest(
			newFuncWithReadRepairFunc(readRepairCaptureQuery(&query)),
			coordinator.GetValueRequest{Key: "abc"},
		)

		assert.Equal(tt, "abc", query)
	})

	t.Run("should call read repair with same context", func(tt *testing.T) {
		var ctx context.Context

		getValueWithContext(
			newFuncWithReadRepairFunc(readRepairCaptureContext(&ctx)),
			cntx.WithValue("name", "same ctx"),
		)

		assert.Equal(tt, "same ctx", ctx.Value("name"))
	})

	t.Run("should return from read repair", func(tt *testing.T) {
		resp := getValueWithResponse(
			newFuncWithReadRepairFunc(readRepairWithResponse(result.Value("yeah"))),
		)

		assert.Equal(tt, result.Value(coordinator.GetValueResponse{Value: "yeah"}), resp)
	})
}
