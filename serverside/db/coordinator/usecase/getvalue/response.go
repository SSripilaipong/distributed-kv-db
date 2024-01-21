package getvalue

import "distributed-kv-db/serverside/db/coordinator"

func responseFromValue(value string) coordinator.GetValueResponse {
	return coordinator.GetValueResponse{Value: value}
}
