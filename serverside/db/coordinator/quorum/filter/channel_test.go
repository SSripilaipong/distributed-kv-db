package filter

import (
	"distributed-kv-db/common/rslt"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ChannelToSlice(t *testing.T) {
	type T = int
	channelToSlice := ChannelToSlice[T]

	t.Run("should return error if n is less than 1", func(tt *testing.T) {
		assert.Equal(tt, rslt.Error[[]T](errors.New("n must be more than 0")), channelToSlice(0)(nil))
	})
}
