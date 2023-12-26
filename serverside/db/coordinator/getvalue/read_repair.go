package getvalue

import "distributed-kv-db/serverside/db/coordinator/quorum"

type readRepairFunc = quorum.ReadRepairFunc[string, string]

var newReadRepairFunc = quorum.NewReadRepairFunc[string, string]
