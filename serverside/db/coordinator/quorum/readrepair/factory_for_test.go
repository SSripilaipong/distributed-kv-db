package readrepair

import (
	"distributed-kv-db/serverside/db/coordinator/quorum"
)

func newFuncWithQuorumRead[Key, Data any](quorumRead quorum.ReadFunc[Key, Data]) quorum.ReadFunc[Key, Data] {
	return newFunc(quorumRead, quorumWrite[Key, Data]())
}

func newFuncWithQuorumWrite[Key, Data any](quorumWrite quorum.WriteFunc[Key, Data]) quorum.ReadFunc[Key, Data] {
	return newFunc(quorumRead[Key, Data](), quorumWrite)
}

func newFuncWithQuorumReadAndQuorumWrite[Key, Data any](quorumRead quorum.ReadFunc[Key, Data], quorumWrite quorum.WriteFunc[Key, Data]) quorum.ReadFunc[Key, Data] {
	return newFunc(quorumRead, quorumWrite)
}
