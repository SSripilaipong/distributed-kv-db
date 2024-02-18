package temporal

import (
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/common/slc"
	"errors"
)

func LatestInSlice[Data Hashable](xs []Data) rslt.Of[Data] {
	if len(xs) == 0 {
		return rslt.Error[Data](errors.New("no data"))
	}
	return rslt.Value(value(slc.SafeMaxComparables(slc.Map(NewWithHashComparison[Data], xs))))
}

type valuable[T any] interface {
	Value() T
}

func value[T valuable[V], V any](x T) V {
	return x.Value()
}
