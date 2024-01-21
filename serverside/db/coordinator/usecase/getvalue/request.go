package getvalue

import "distributed-kv-db/serverside/db/coordinator"

func keyOfRequest(request coordinator.GetValueRequest) string {
	return request.Key
}
