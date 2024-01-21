package write

import (
	"distributed-kv-db/serverside/db/coordinator/peer/discovery"
	"distributed-kv-db/serverside/db/coordinator/peer/read"
)

func New[Key, Data any](_ discovery.Func[Key, read.ReadableNode[Key, Data]]) Func[Key, Data] {
	return nil
}
