package getvalue

import (
	"context"
	"distributed-kv-db/common/cntx"
	"distributed-kv-db/serverside/db/coordinator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_newFunc(t *testing.T) {
	t.Run("should call read repair with key as query", func(tt *testing.T) {
		var query string

		getValueWithRequest(
			newFunc(readRepairCaptureQuery(&query)),
			coordinator.GetValueRequest{Key: "abc"},
		)

		assert.Equal(tt, "abc", query)
	})

	t.Run("should call read repair with same context", func(tt *testing.T) {
		var ctx context.Context

		getValueWithContext(
			newFunc(readRepairCaptureContext(&ctx)),
			cntx.WithValue("name", "same ctx"),
		)

		assert.Equal(tt, "same ctx", ctx.Value("name"))
	})
}
