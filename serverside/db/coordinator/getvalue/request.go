package getvalue

import "distributed-kv-db/serverside/db/coordinator"

func requestToKey(request coordinator.GetValueRequest) string {
	return request.Key
}
