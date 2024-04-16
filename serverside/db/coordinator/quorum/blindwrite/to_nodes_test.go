package blindwrite

import (
	"context"
	"distributed-kv-db/common/chn"
	"distributed-kv-db/common/cntx"
	"distributed-kv-db/common/tstexc"
	"distributed-kv-db/common/tstmck"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ToNodes(t *testing.T) {
	type Key = int
	type Data = int
	type Node = string
	Target := tstexc.NewFunc2Dep4In1Out(composeToNodes[Key, Data, Node])
	QFilter := tstmck.NewFunc1In1OutLike(Target.D1)
	PBlindWrite := tstmck.NewFunc4In1OutLike(Target.D2)
	Target = Target.
		DefaultDep1(QFilter.ReturnO1(nil)).
		DefaultDep2(PBlindWrite.ReturnO1(chn.Closed(make(chan error))))
	WithQFilter := Target.WithDep1
	WithPBlindWrite := Target.WithDep2
	WithContext := Target.WithIn1
	WithNodes := Target.WithIn2
	WithKey := Target.WithIn3
	WithData := Target.WithIn4
	PBlindWriteCaptureContext := PBlindWrite.CaptureI1
	PBlindWriteCaptureKey := PBlindWrite.CaptureI2
	PBlindWriteCaptureData := PBlindWrite.CaptureI3
	PBlindWriteCaptureNodes := PBlindWrite.CaptureI4
	PBlindWriteWithResult := PBlindWrite.ReturnO1
	QFilterCaptureChannel := QFilter.CaptureI1
	QFilterWithResult := QFilter.ReturnO1

	t.Run("should write to nodes", func(t *testing.T) {
		var capturedNodes []Node
		_ = Target.Execute(
			WithNodes([]Node{"aaa", "bbb"}),
			WithPBlindWrite(PBlindWriteCaptureNodes(&capturedNodes)),
		)
		assert.ElementsMatch(t, []Node{"aaa", "bbb"}, capturedNodes)
	})

	t.Run("should write with same context", func(t *testing.T) {
		var capturedCtx context.Context
		ctx, isSame := cntx.WithVerifier()
		_ = Target.Execute(
			WithContext(ctx),
			WithPBlindWrite(PBlindWriteCaptureContext(&capturedCtx)),
		)
		assert.True(t, isSame(capturedCtx))
	})

	t.Run("should write with key", func(t *testing.T) {
		var capturedKey Key
		_ = Target.Execute(
			WithKey(999),
			WithPBlindWrite(PBlindWriteCaptureKey(&capturedKey)),
		)
		assert.Equal(t, 999, capturedKey)
	})

	t.Run("should write with data", func(t *testing.T) {
		var capturedData Data
		_ = Target.Execute(
			WithData(999),
			WithPBlindWrite(PBlindWriteCaptureData(&capturedData)),
		)
		assert.Equal(t, 999, capturedData)
	})

	t.Run("should filter write error", func(t *testing.T) {
		ch := make(chan error)
		var capturedChannel <-chan error
		_ = Target.Execute(
			WithPBlindWrite(PBlindWriteWithResult(ch)),
			WithQFilter(QFilterCaptureChannel(&capturedChannel)),
		)
		assert.Equal(t, (<-chan error)(ch), capturedChannel)
	})

	t.Run("should return error from filter", func(t *testing.T) {
		err := Target.Execute(
			WithQFilter(QFilterWithResult(errors.New("boom"))),
		)

		assert.Equal(t, errors.New("boom"), err)
	})
}
