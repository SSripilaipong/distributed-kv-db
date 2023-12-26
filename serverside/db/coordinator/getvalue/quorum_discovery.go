package getvalue

import "distributed-kv-db/serverside/db/coordinator/quorum"

type quorumDiscovery quorum.Discovery[string, string]
