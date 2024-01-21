package request

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NodesToChannel(t *testing.T) {
	type Node = string
	type Data = int
	NewTarget := nodesToChannelForTest[Node, Data]
	Execute := runNodesToChannelForTest[Node, Data]

	t.Run("should return channel", func(tt *testing.T) {
		assert.NotNil(tt, Execute(NewTarget()))
	})
}
