package readrepair

import (
	"distributed-kv-db/serverside/db/coordinator/quorum"
)

func newFuncWithReadRepair[Key, Data any](quorumRead quorum.ReadFunc[Key, Data]) quorum.ReadFunc[Key, Data] {
	return newFunc(quorumRead)
}
