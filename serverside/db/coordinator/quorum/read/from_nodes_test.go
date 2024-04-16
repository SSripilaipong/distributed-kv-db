package read

import (
	"context"
	"distributed-kv-db/common/chn"
	"distributed-kv-db/common/cntx"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/common/tstexc"
	"distributed-kv-db/common/tstmck"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_FromNodes(t *testing.T) {
	type Key = int
	type Data = int
	type Node = string
	Target := tstexc.NewFunc2Dep3In1Out(composeFromNodes[Key, Data, Node])
	QFilter := tstmck.NewFunc1In1OutLike(Target.D1)
	PRead := tstmck.NewFunc3In1OutLike(Target.D2)
	Target = Target.
		DefaultDep1(QFilter.ReturnO1(rslt.Value([]Data{}))).
		DefaultDep2(PRead.ReturnO1(chn.Closed(make(chan Data))))
	WithQFilter := Target.WithDep1
	WithPRead := Target.WithDep2
	WithContext := Target.WithIn1
	WithNodes := Target.WithIn2
	WithKey := Target.WithIn3
	PReadCaptureContext := PRead.CaptureI1
	PReadCaptureKey := PRead.CaptureI2
	PReadCaptureNodes := PRead.CaptureI3
	PReadWithResult := PRead.ReturnO1
	QFilterCaptureChannel := QFilter.CaptureI1
	QFilterWithResult := QFilter.ReturnO1

	t.Run("should read from nodes", func(t *testing.T) {
		var capturedNodes []Node
		Target.Execute(
			WithPRead(PReadCaptureNodes(&capturedNodes)),
			WithNodes([]Node{"a", "b"}),
		)
		assert.Equal(t, []Node{"a", "b"}, capturedNodes)
	})

	t.Run("should read with context", func(t *testing.T) {
		var capturedContext context.Context
		ctx, isSame := cntx.WithVerifier()
		Target.Execute(
			WithPRead(PReadCaptureContext(&capturedContext)),
			WithContext(ctx),
		)
		assert.True(t, isSame(capturedContext))
	})

	t.Run("should read with key", func(t *testing.T) {
		var capturedKey Key
		Target.Execute(
			WithPRead(PReadCaptureKey(&capturedKey)),
			WithKey(999),
		)
		assert.Equal(t, 999, capturedKey)
	})

	t.Run("should filter read channel", func(t *testing.T) {
		ch := make(chan Data)
		var capturedChannel <-chan Data
		Target.Execute(
			WithPRead(PReadWithResult(ch)),
			WithQFilter(QFilterCaptureChannel(&capturedChannel)),
		)
		assert.Equal(t, (<-chan Data)(ch), capturedChannel)
	})

	t.Run("should return filter result", func(t *testing.T) {
		result := Target.Execute(
			WithQFilter(QFilterWithResult(rslt.Value([]Data{123, 456}))),
		)
		assert.Equal(t, rslt.Value([]Data{123, 456}), result)
	})
}
