package filter

import (
	"distributed-kv-db/common/rslt"
	"errors"
)

func ChannelToSlice[T any](n int) func(<-chan T) rslt.Of[[]T] {
	return func(ch <-chan T) rslt.Of[[]T] {
		return rslt.Error[[]T](errors.New("n must be more than 0"))
	}
}
