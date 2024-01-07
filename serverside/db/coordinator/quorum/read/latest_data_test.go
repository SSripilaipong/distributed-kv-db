package read

import (
	"distributed-kv-db/common/rslt"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_latestData(t *testing.T) {
	t.Run("should return error if slice is empty", func(tt *testing.T) {
		result := latestData[orderableDataMock]([]orderableDataMock{})
		assert.Equal(tt, rslt.Error[orderableDataMock](errors.New("no data")), result)
	})

	t.Run("should return error if slice is nil", func(tt *testing.T) {
		result := latestData[orderableDataMock](nil)
		assert.Equal(tt, rslt.Error[orderableDataMock](errors.New("no data")), result)
	})
}

type orderableDataMock struct {
}

func (d orderableDataMock) Newness() int {
	//TODO implement me
	panic("implement me")
}

func (d orderableDataMock) Hash() string {
	//TODO implement me
	panic("implement me")
}
