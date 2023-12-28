package write

import "distributed-kv-db/serverside/db/coordinator/quorum"

func New[Key, Data any](_ quorum.Discovery[Key, Data]) quorum.WriteFunc[Key, Data] {
	return nil
}
