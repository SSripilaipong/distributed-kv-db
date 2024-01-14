package filter

import (
	"distributed-kv-db/common/chn"
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/common/zd"
	"errors"
)

func ChannelToSlice[T any](n int) func(<-chan T) rslt.Of[[]T] {
	nOrErr := mustBeMoreThan(0)(n)
	return rslt.MapExecutePartial(rslt.Fmap(fn.Compose(chn.FirstNFunc[T], numberOfQuorum))(nOrErr))
}

var numberOfQuorum = fn.Compose(zd.Successor, zd.Half)

func mustBeMoreThan(m int) func(n int) rslt.Of[int] {
	return func(n int) rslt.Of[int] {
		if n <= m {
			return rslt.Error[int](errors.New("n must be more than 0"))
		}
		return rslt.Value(n)
	}
}
