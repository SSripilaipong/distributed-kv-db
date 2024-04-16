package temporal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Newer(t *testing.T) {
	t.Run("should return newer instance", func(t *testing.T) {
		newer := dataMockWithNewness(9)
		older := dataMockWithNewness(5)

		assert.Equal(t, newer, Newer(newer, older))
		assert.Equal(t, newer, Newer(older, newer))
	})

	t.Run("should return one with greater hash value if their newness equal", func(t *testing.T) {
		greater := dataMockWithNewnessAndHash(4, 999)
		lesser := dataMockWithNewnessAndHash(4, 111)

		assert.Equal(t, greater, Newer(greater, lesser))
		assert.Equal(t, greater, Newer(lesser, greater))
	})

	t.Run("should return any of them if both newness and hash equal", func(t *testing.T) {
		a := dataMockWithNewnessAndHash(4, 999)
		b := dataMockWithNewnessAndHash(4, 999)

		assert.True(t, a == Newer(a, b))
		assert.True(t, b == Newer(a, b))
	})
}
