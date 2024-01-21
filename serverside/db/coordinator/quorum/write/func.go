package write

import (
	"context"
)

type Func[Key, Data any] func(context.Context, Key, Data) error
