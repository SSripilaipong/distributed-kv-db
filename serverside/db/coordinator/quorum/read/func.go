package read

import (
	"context"
	"distributed-kv-db/common/rslt"
)

type Func[Key, Data any] func(context.Context, Key) rslt.Of[Data]
