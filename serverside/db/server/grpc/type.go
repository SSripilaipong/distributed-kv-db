package grpc

import "distributed-kv-db/common/result"

type Func func(port int) result.Of[RunningServer]
