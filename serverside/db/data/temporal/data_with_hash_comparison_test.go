package temporal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_DataWithHashComparison_Compare(t *testing.T) {
	t.Run("should return 0 if hash and newness equals", func(tt *testing.T) {
		a := newHashableDataForTest(1, 2)
		b := newHashableDataForTest(1, 2)
		assert.Equal(tt, 0, a.Compare(b))
		assert.Equal(tt, 0, b.Compare(a))
	})

	t.Run("should return negative if newness is less than other", func(tt *testing.T) {
		a := newHashableDataForTest(1, 2)
		b := newHashableDataForTest(3, 2)
		assert.Negative(tt, a.Compare(b))
	})

	t.Run("should return positive if newness is greater than other", func(tt *testing.T) {
		a := newHashableDataForTest(3, 2)
		b := newHashableDataForTest(1, 2)
		assert.Positive(tt, a.Compare(b))
	})

	t.Run("should return negative if newness equals but hash is less than other", func(tt *testing.T) {
		a := newHashableDataForTest(3, 2)
		b := newHashableDataForTest(3, 5)
		assert.Negative(tt, a.Compare(b))
	})

	t.Run("should return negative if newness equals but hash is greater than other", func(tt *testing.T) {
		a := newHashableDataForTest(3, 7)
		b := newHashableDataForTest(3, 5)
		assert.Positive(tt, a.Compare(b))
	})
}
