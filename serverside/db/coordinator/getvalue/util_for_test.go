package getvalue

import (
	"context"
	"distributed-kv-db/common/result"
)

func readRepairCaptureQuery(query *string) readRepairFunc {
	return func(ctx context.Context, q string) result.Of[string] {
		*query = q
		return result.Value("")
	}
}
