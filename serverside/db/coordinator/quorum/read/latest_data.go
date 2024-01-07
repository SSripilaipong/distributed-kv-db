package read

import (
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/serverside/db/coordinator/quorum"
	"errors"
	"slices"
)

type orderableData interface {
	quorum.Orderable
	quorum.Hashable
}

func latestData[Data orderableData](xs []Data) rslt.Of[Data] {
	if len(xs) == 0 {
		return rslt.Error[Data](errors.New("no data"))
	}
	return rslt.Value(slices.MaxFunc(xs, func(x, y Data) int {
		return x.Newness() - y.Newness()
	}))
}
