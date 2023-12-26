package getvalue

import "distributed-kv-db/serverside/db/coordinator"

func valueToResponse(value string) coordinator.GetValueResponse {
	return coordinator.GetValueResponse{Value: value}
}
