package filter

import (
	"distributed-kv-db/common/chn"
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/common/zd"
)

func ChannelToSlice[T any](n uint) func(<-chan T) rslt.Of[[]T] {
	filter := fn.Compose(chn.FirstNFunc[T], nToQuorum)
	filterOrError := fn.Compose(rslt.Fmap(filter), zd.MustBeMoreThan(0))
	return rslt.MapOfFuncPartial(filterOrError(int(n)))
}
