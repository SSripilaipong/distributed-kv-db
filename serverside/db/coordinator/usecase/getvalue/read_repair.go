package getvalue

import (
	"distributed-kv-db/serverside/db/coordinator/quorum/read"
)

type readRepairFunc = read.Func[string, string]
