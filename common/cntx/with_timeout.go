package cntx

import (
	"context"
	"time"
)

func WithTimeout(duration time.Duration) context.Context {
	ctx, _ := context.WithTimeout(context.Background(), duration)
	return ctx
}
