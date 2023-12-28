package read

import (
	"context"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/common/typ"
	"distributed-kv-db/serverside/db/coordinator/quorum"
)

func read[Key, Data any](read quorum.ReadFunc[Key, Data]) rslt.Of[Data] {
	return read(context.Background(), typ.Zero[Key]())
}
