package write

import "distributed-kv-db/serverside/db/coordinator/quorum"

func New[Key, Data any](_ quorum.DiscoverNodes[Key, quorum.ReadNode[Key, Data]]) quorum.WriteFunc[Key, Data] {
	return nil
}
