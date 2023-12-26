package getvalue

import (
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
}
