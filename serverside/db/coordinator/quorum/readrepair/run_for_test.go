package readrepair

import (
	"context"
	"distributed-kv-db/common/result"
	"distributed-kv-db/serverside/db/coordinator/quorum"
)

func readRepairWithKey[Key, Data any](readRepair quorum.ReadFunc[Key, Data], key Key) result.Of[Data] {
	return readRepair(context.Background(), key)
}
