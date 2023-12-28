package fn

import "context"

func Ctx[A, B any](ctx context.Context, f func(ctx context.Context, x A) (y B)) func(x A) (y B) {
	return Bind(ctx, f)
}

func Ctx2[A1, A2, B any](ctx context.Context, f func(ctx context.Context, x1 A1, x2 A2) (y B)) func(x1 A1, x2 A2) (y B) {
	return func(x1 A1, x2 A2) B {
		return f(ctx, x1, x2)
	}
}
