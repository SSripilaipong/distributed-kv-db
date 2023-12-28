package read

import (
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/serverside/db/coordinator/quorum"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_New(t *testing.T) {

	t.Run("should read nodes from discovery", func(tt *testing.T) {
		var nodes []quorum.Node[int, int]
		read(newFunc[int, int](
			discoveryWithNodes(rslt.Value([]quorum.Node[int, int]{nodeWithId[int, int]{1}, nodeWithId[int, int]{2}})),
			readNodesToChannelCaptureNodes(&nodes),
		))
		assert.Equal(tt, []quorum.Node[int, int]{nodeWithId[int, int]{1}, nodeWithId[int, int]{2}}, nodes)
	})

}
