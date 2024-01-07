package chn

import (
	"context"
	"errors"
)

func SendToWithContext[T any](ctx context.Context, ch chan<- T) func(x T) error {
	return func(x T) error {
		select {
		case ch <- x:
			return nil
		case <-ctx.Done():
			return errors.New("context is done")
		}
	}
}
