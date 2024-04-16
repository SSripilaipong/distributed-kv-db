package readrepair

import (
	peerBlindWrite "distributed-kv-db/serverside/db/coordinator/peer/blindwrite"
	peerRead "distributed-kv-db/serverside/db/coordinator/peer/read"
)

type ReadWritableNode[Key, Data any] interface {
	peerRead.ReadableNode[Key, Data]
	peerBlindWrite.WritableNode[Key, Data]
}
