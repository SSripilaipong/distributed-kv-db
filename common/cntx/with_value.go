package cntx

import "context"

func WithValue(key, value any) context.Context {
	return context.WithValue(context.Background(), key, value)
}
