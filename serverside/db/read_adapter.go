package db

import "distributed-kv-db/serverside/db/coordinator/quorum"

func readAdapter(read quorum.ReadFunc[string, orderableDataAdapter]) quorum.ReadFunc[string, string] {
	//TODO implement me
	panic("implement me")
}
