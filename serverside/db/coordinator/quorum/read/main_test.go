package read

import (
	"context"
	"distributed-kv-db/common/chn"
	"distributed-kv-db/common/cntx"
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/rslt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_NodesToDataSlice(t *testing.T) {
	type Key int
	type Data int
	type Node string
	defaultTimeout := 100 * time.Millisecond
	Target := newNodesToDataSliceForTest[Key, Data, Node]
	WithReadNodes := withReadNodes[Key, Data, Node]
	WithDiscoverNodes := withDiscoverNodes[Key, Data, Node]
	WithFilterQuorum := withFilterQuorum[Key, Data, Node]
	DiscoverNodesCaptureKey := discoverNodesCaptureKey[Key, Node]
	DiscoverNodesCaptureContext := discoverNodesCaptureContext[Key, Node]
	DiscoverNodesWithResult := discoverNodesWithResult[Key, Node]
	ReadNodesCaptureNodes := readNodesCaptureNodes[Key, Data, Node]
	ReadNodesCaptureKey := readNodesCaptureKey[Key, Data, Node]
	ReadNodesCaptureContext := readNodesCaptureContext[Key, Data, Node]
	ReadNodesWithResult := readNodesWithResult[Key, Data, Node]
	FilterQuorumCaptureXs := filterQuorumCaptureXs[Data]
	FilterQuorumCaptureN := filterQuorumCaptureN[Data]
	Execute := runNodesToDataSliceForTest[Key, Data, Node]
	WithKey := withKey[Key, Node]
	WithContext := withContext[Key, Node]
	AllData := fn.WithArg(cntx.WithTimeout(defaultTimeout), chn.AllWithCtx[Data])

	t.Run("should discover nodes with key", func(tt *testing.T) {
		var key Key
		Execute(
			Target(WithDiscoverNodes(DiscoverNodesCaptureKey(&key))),
			WithKey(123),
		)
		assert.Equal(tt, Key(123), key)
	})

	t.Run("should discover nodes with same context", func(tt *testing.T) {
		inputCtx, isSame := cntx.WithVerifier()

		var ctx context.Context
		Execute(
			Target(WithDiscoverNodes(DiscoverNodesCaptureContext(&ctx))),
			WithContext(inputCtx),
		)
		assert.True(tt, isSame(ctx))
	})

	t.Run("should discover nodes with same context", func(tt *testing.T) {
		inputCtx, isSame := cntx.WithVerifier()

		var ctx context.Context
		Execute(
			Target(WithDiscoverNodes(DiscoverNodesCaptureContext(&ctx))),
			WithContext(inputCtx),
		)
		assert.True(tt, isSame(ctx))
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

	t.Run("should read nodes with context", func(tt *testing.T) {
		inputCtx, isSame := cntx.WithVerifier()

		var ctx context.Context
		Execute(
			Target(WithReadNodes(ReadNodesCaptureContext(&ctx))),
			WithContext(inputCtx),
		)
		assert.True(tt, isSame(ctx))
	})

	t.Run("should filter quorum from read not channel", func(tt *testing.T) {
		var xs <-chan Data
		Execute(Target(
			WithReadNodes(ReadNodesWithResult(chn.NewFromSlice([]Data{1, 2, 3}))),
			WithFilterQuorum(FilterQuorumCaptureXs(&xs)),
		))
		assert.Equal(tt, rslt.Value([]Data{1, 2, 3}), AllData(xs))
	})

	t.Run("should filter quorum with number of nodes", func(tt *testing.T) {
		var n int
		Execute(Target(
			WithDiscoverNodes(DiscoverNodesWithResult(rslt.Value([]Node{"a", "b", "c"}))),
			WithFilterQuorum(FilterQuorumCaptureN(&n)),
		))
		assert.Equal(tt, 3, n)
	})

	t.Run("should return from filter quorum", func(tt *testing.T) {
		y := Execute(Target(
			WithFilterQuorum(filterQuorumWithResult(rslt.Value([]Data{3, 2, 1}))),
		))
		assert.Equal(tt, rslt.Value([]Data{3, 2, 1}), y)
	})
}
