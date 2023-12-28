package readrepair

import (
	"context"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/serverside/db/coordinator/quorum"
)

func readRepairWithKey[Key, Data any](readRepair quorum.ReadFunc[Key, Data], key Key) rslt.Of[Data] {
	return readRepair(context.Background(), key)
}
