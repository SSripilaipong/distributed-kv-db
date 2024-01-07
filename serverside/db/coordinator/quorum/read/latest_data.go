package read

import (
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/serverside/db/coordinator/quorum"
	"errors"
)

type orderableData interface {
	quorum.Orderable
	quorum.Hashable
}

func latestData[Data orderableData](_ []Data) rslt.Of[Data] {
	return rslt.Error[Data](errors.New("no data"))
}
