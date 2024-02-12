package cntx

import (
	"context"
	"github.com/google/uuid"
)

func WithVerifier() (ctx context.Context, isSame func(context.Context) bool) {
	key := "__cntx.WithVerification"
	id := uuid.NewString()
	return context.WithValue(context.Background(), key, id), func(ctx context.Context) bool {
		if ctx == nil {
			return false
		}
		return ctx.Value(key) == id
	}
}
