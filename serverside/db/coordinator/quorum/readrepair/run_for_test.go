package readrepair

import (
	"context"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/common/typ"
	"distributed-kv-db/serverside/db/coordinator/quorum"
)

func readRepair[Key, Data any](readRepair quorum.ReadFunc[Key, Data]) rslt.Of[Data] {
	return readRepair(context.Background(), typ.Zero[Key]())
}

func readRepairWithContext[Key, Data any](readRepair quorum.ReadFunc[Key, Data], ctx context.Context) rslt.Of[Data] {
	return readRepair(ctx, typ.Zero[Key]())
}

func readRepairWithKey[Key, Data any](readRepair quorum.ReadFunc[Key, Data], key Key) rslt.Of[Data] {
	return readRepair(context.Background(), key)
}
