package read

import (
	"context"
	"distributed-kv-db/common/cntx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NodesToDataSlice(t *testing.T) {
	type Key int
	type Data int
	type Node string
	NewTarget := newNodesToDataSliceForTest[Key, Data, Node]
	WithDiscoverNodes := withDiscoverNodes[Key, Data, Node]
	DiscoverNodesCaptureKey := discoverNodesCaptureKey[Key, Node]
	DiscoverNodesCaptureContext := discoverNodesCaptureContext[Key, Node]
	Execute := runNodesToDataSliceForTest[Key, Data, Node]
	WithKey := withKey[Key, Node]
	WithContext := withContext[Key, Node]

	t.Run("should discover nodes with key", func(tt *testing.T) {
		var key Key
		Execute(
			NewTarget(WithDiscoverNodes(DiscoverNodesCaptureKey(&key))),
			WithKey(123),
		)
		assert.Equal(tt, Key(123), key)
	})

	t.Run("should discover nodes with same context", func(tt *testing.T) {
		var ctx context.Context
		Execute(
			NewTarget(WithDiscoverNodes(DiscoverNodesCaptureContext(&ctx))),
			WithContext(cntx.WithValue("is the same?", "yes")),
		)
		assert.Equal(tt, "yes", ctx.Value("is the same?"))
	})
}
