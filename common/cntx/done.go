package cntx

import (
	"context"
)

func Done() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	return ctx
}
