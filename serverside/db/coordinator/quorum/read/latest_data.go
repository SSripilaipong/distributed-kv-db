package read

import (
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/common/typ"
	"distributed-kv-db/serverside/db/coordinator/quorum"
)

type orderableData interface {
	quorum.Orderable
	quorum.Hashable
}

func latestData[Data orderableData](_ []Data) rslt.Of[Data] {
	return rslt.Value(typ.Zero[Data]())
}
