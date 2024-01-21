package read

import (
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
	Execute := runNodesToDataSliceForTest[Key, Data, Node]
	WithKey := withKey[Key, Node]

	t.Run("should discover nodes with key", func(tt *testing.T) {
		var key Key
		Execute(
			NewTarget(WithDiscoverNodes(DiscoverNodesCaptureKey(&key))),
			WithKey(123),
		)
		assert.Equal(tt, Key(123), key)
	})
}
