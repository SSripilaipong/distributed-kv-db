package fn

import "context"

func Ctx[A, B any](ctx context.Context, f func(ctx context.Context, x A) (y B)) func(x A) (y B) {
	return func(x A) (y B) {
		return f(ctx, x)
	}
}
