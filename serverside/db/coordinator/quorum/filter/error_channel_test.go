package filter

import (
	"distributed-kv-db/common/chn"
	"distributed-kv-db/common/tstexc"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ErrorChannel(t *testing.T) {
	Target := tstexc.NewFunc1Dep1In1Out(ErrorChannel)
	WithN := Target.WithDep1
	WithChannel := Target.WithIn1

	t.Run("should return error for closed channel", func(t *testing.T) {
		err := Target.Execute(
			WithChannel(chn.Closed(make(chan error))),
		)

		assert.Equal(t, err.Error(), "channel closed")
	})

	t.Run("should return no error when all errors are nil", func(t *testing.T) {
		err := Target.Execute(
			WithN(2),
			WithChannel(chn.NewFromSlice([]error{nil, nil})),
		)

		assert.Nil(t, err)
	})

	t.Run("should return no error when a quorum of errors are nil", func(t *testing.T) {
		err := Target.Execute(
			WithN(3),
			WithChannel(chn.NewFromSlice([]error{nil, errors.New(":D"), nil})),
		)

		assert.Nil(t, err)
	})

	t.Run("should return joined error", func(t *testing.T) {
		err := Target.Execute(
			WithN(3),
			WithChannel(chn.NewFromSlice([]error{nil, errors.New("failed haha"), errors.New("boom")})),
		)

		assert.Contains(t, err.Error(), "boom")
		assert.Contains(t, err.Error(), "failed haha")
	})

	t.Run("should return the error joined with channel closed", func(t *testing.T) {
		err := Target.Execute(
			WithN(3),
			WithChannel(chn.NewFromSlice([]error{nil, errors.New("boom")})),
		)

		assert.Contains(t, err.Error(), "boom")
		assert.Contains(t, err.Error(), "channel closed")
	})
}
