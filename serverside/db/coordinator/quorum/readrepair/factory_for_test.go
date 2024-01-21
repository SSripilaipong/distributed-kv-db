package readrepair

import (
	"distributed-kv-db/serverside/db/coordinator/quorum/read"
	"distributed-kv-db/serverside/db/coordinator/quorum/write"
)

func newFuncWithQuorumRead[Key, Data any](quorumRead read.Func[Key, Data]) read.Func[Key, Data] {
	return newFunc(quorumRead, quorumWrite[Key, Data]())
}

func newFuncWithQuorumWrite[Key, Data any](quorumWrite write.Func[Key, Data]) read.Func[Key, Data] {
	return newFunc(quorumRead[Key, Data](), quorumWrite)
}

func newFuncWithQuorumReadAndQuorumWrite[Key, Data any](quorumRead read.Func[Key, Data], quorumWrite write.Func[Key, Data]) read.Func[Key, Data] {
	return newFunc(quorumRead, quorumWrite)
}
