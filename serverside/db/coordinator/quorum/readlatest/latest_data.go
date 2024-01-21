package readlatest

import (
	"distributed-kv-db/common/rslt"
	"errors"
	"slices"
)

type ReadableData interface {
	Orderable
	Hashable
}

func latestData[Data ReadableData](xs []Data) rslt.Of[Data] {
	if len(xs) == 0 {
		return rslt.Error[Data](errors.New("no data"))
	}
	return rslt.Value(slices.MaxFunc(xs, func(x, y Data) int {
		if newnessDiff := x.Newness() - y.Newness(); newnessDiff != 0 {
			return newnessDiff
		}
		return x.Hash() - y.Newness()
	}))
}

type Orderable interface {
	Newness() int
}

type Hashable interface {
	Hash() int
}
