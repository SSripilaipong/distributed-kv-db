package read

import (
	"context"
	"distributed-kv-db/common/cntx"
	"distributed-kv-db/common/rslt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NodesToDataSlice(t *testing.T) {
	type Key int
	type Data int
	type Node string
	Target := newNodesToDataSliceForTest[Key, Data, Node]
	WithReadNodes := withReadNodes[Key, Data, Node]
	WithDiscoverNodes := withDiscoverNodes[Key, Data, Node]
	DiscoverNodesCaptureKey := discoverNodesCaptureKey[Key, Node]
	DiscoverNodesCaptureContext := discoverNodesCaptureContext[Key, Node]
	DiscoverNodesWithResult := discoverNodesWithResult[Key, Node]
	ReadNodesCaptureNodes := readNodesCaptureNodes[Key, Data, Node]
	ReadNodesCaptureKey := readNodesCaptureKey[Key, Data, Node]
	Execute := runNodesToDataSliceForTest[Key, Data, Node]
	WithKey := withKey[Key, Node]
	WithContext := withContext[Key, Node]

	t.Run("should discover nodes with key", func(tt *testing.T) {
		var key Key
		Execute(
			Target(WithDiscoverNodes(DiscoverNodesCaptureKey(&key))),
			WithKey(123),
		)
		assert.Equal(tt, Key(123), key)
	})

	t.Run("should discover nodes with same context", func(tt *testing.T) {
		var ctx context.Context
		Execute(
			Target(WithDiscoverNodes(DiscoverNodesCaptureContext(&ctx))),
			WithContext(cntx.WithValue("is the same?", "yes")),
		)
		assert.Equal(tt, "yes", ctx.Value("is the same?"))
	})

	t.Run("should discover nodes with same context", func(tt *testing.T) {
		var ctx context.Context
		Execute(
			Target(WithDiscoverNodes(DiscoverNodesCaptureContext(&ctx))),
			WithContext(cntx.WithValue("is the same?", "yes")),
		)
		assert.Equal(tt, "yes", ctx.Value("is the same?"))
	})

	t.Run("should read from discovered nodes", func(tt *testing.T) {
		var nodes []Node
		Execute(Target(
			WithDiscoverNodes(DiscoverNodesWithResult(rslt.Value([]Node{"node1", "node2"}))),
			WithReadNodes(ReadNodesCaptureNodes(&nodes)),
		))
		assert.Equal(tt, []Node{"node1", "node2"}, nodes)
	})

	t.Run("should read nodes with key", func(tt *testing.T) {
		var key Key
		Execute(
			Target(WithReadNodes(ReadNodesCaptureKey(&key))),
			WithKey(555),
		)
		assert.Equal(tt, Key(555), key)
	})
}
