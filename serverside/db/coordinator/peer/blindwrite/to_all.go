package blindwrite

import (
	"context"
	"distributed-kv-db/common/rslt"
)

func ToAll[Key, Data, Node any](
	discoverNodes discoverFunc[Key, Node],
) func(context.Context, Data) <-chan error {
	return composeToAll[Key, Data, Node](discoverNodes)
}

func composeToAll[Key, Data, Node any](
	discoverNodes discoverFunc[Key, Node],
) func(context.Context, Data) <-chan error {
	return func(ctx context.Context, data Data) <-chan error {
		return nil
	}
}

type discoverFunc[Key, Node any] func(ctx context.Context, key Key) rslt.Of[[]Node]
