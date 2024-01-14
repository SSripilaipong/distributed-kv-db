package filter

import (
	"distributed-kv-db/common/chn"
	"distributed-kv-db/common/rslt"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ChannelToSlice(t *testing.T) {
	type T = int
	channelToSlice := ChannelToSlice[T]

	t.Run("should return error if n is less than 1", func(tt *testing.T) {
		assert.Equal(tt, rslt.Error[[]T](errors.New("n must be more than 0")), channelToSlice(0)(chn.NewFromSlice([]T{})))
	})

	t.Run("should return first element if n is 1", func(tt *testing.T) {
		result := channelToSlice(1)(chn.NewFromSlice([]T{1, 2, 3}))
		assert.Equal(tt, rslt.Value([]T{1}), result)
	})

	t.Run("should return first 2 elements when n is 3", func(tt *testing.T) {
		result := channelToSlice(3)(chn.NewFromSlice([]T{1, 2, 3}))
		assert.Equal(tt, rslt.Value([]T{1, 2}), result)
	})
}
