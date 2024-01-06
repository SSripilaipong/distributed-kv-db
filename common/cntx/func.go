package cntx

import "context"

func Func[A, B any](f func(A) B) func(context.Context, A) B {
	return func(ctx context.Context, x A) B {
		return f(x)
	}
}

func Func2[A1, A2, B any](f func(A1, A2) B) func(context.Context, A1, A2) B {
	return func(ctx context.Context, x1 A1, x2 A2) B {
		return f(x1, x2)
	}
}
