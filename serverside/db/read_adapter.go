package db

import (
	"distributed-kv-db/serverside/db/coordinator/quorum/read"
)

func readAdapter(read read.Func[string, orderableDataAdapter]) read.Func[string, string] {
	//TODO implement me
	panic("implement me")
}
