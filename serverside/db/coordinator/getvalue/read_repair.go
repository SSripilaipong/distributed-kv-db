package getvalue

import "distributed-kv-db/serverside/db/coordinator/quorum"

type readRepairFunc = quorum.ReadFunc[string, string]
