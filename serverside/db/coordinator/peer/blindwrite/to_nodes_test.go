package blindwrite

import (
	"context"
	"distributed-kv-db/common/chn"
	"distributed-kv-db/common/cntx"
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/rslt"
	"distributed-kv-db/common/slc"
	"distributed-kv-db/common/tstexc"
	"distributed-kv-db/common/tstmck"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_ToNodes(t *testing.T) {
	type Key = int
	type Data = int
	type Node = string

	defaultTimeout := 500 * time.Millisecond

	Target := tstexc.NewFunc1Dep4In1Out(composeToAll[Key, Data, Node]).
		DefaultIn1(context.Background())
	WithWriteFunc := Target.WithDep1
	WithContext := Target.WithIn1
	WithNodes := Target.WithIn2
	WithKey := Target.WithIn3
	WithData := Target.WithIn4

	WriteFunc := tstmck.NewFunc4In1OutLike(Target.D1)
	WriteFuncCaptureAllContext := WriteFunc.CaptureAllI1
	WriteFuncCaptureAllKey := WriteFunc.CaptureAllI2
	WriteFuncCaptureAllData := WriteFunc.CaptureAllI3
	WriteFuncCaptureAllNode := WriteFunc.CaptureAllI4
	WriteFuncWithResult := WriteFunc.ReturnAllO1

	ReadAllFromChannel := fn.WithArg(cntx.WithTimeout(defaultTimeout), chn.AllWithCtx[error])
	IsChannelClosed := fn.Compose(rslt.IsOk[[]error], ReadAllFromChannel)

	t.Run("should close the channel after all errors are read", func(t *testing.T) {
		assert.True(t, IsChannelClosed(Target.Execute()))
	})

	t.Run("should call write with context", func(t *testing.T) {
		ctx, isSame := cntx.WithVerifier()

		var capturedCtxs []context.Context
		ReadAllFromChannel(Target.Execute(
			WithNodes([]Node{"", ""}),
			WithWriteFunc(WriteFuncCaptureAllContext(&capturedCtxs)),
			WithContext(ctx),
		))
		assert.Equal(t, []bool{true, true}, slc.Map(isSame, capturedCtxs))
	})

	t.Run("should call write with key", func(t *testing.T) {
		var keys []Key
		ReadAllFromChannel(Target.Execute(
			WithNodes([]Node{"", ""}),
			WithKey(111),
			WithWriteFunc(WriteFuncCaptureAllKey(&keys)),
		))
		assert.Equal(t, []Key{111, 111}, keys)
	})

	t.Run("should call write with data", func(t *testing.T) {
		var allData []Data
		ReadAllFromChannel(Target.Execute(
			WithNodes([]Node{"", ""}),
			WithData(999),
			WithWriteFunc(WriteFuncCaptureAllData(&allData)),
		))
		assert.Equal(t, []Key{999, 999}, allData)
	})

	t.Run("should call write for each of the nodes", func(t *testing.T) {
		var nodes []Node
		ReadAllFromChannel(Target.Execute(
			WithWriteFunc(WriteFuncCaptureAllNode(&nodes)),
			WithNodes([]Node{"node1", "node2"}),
		))
		assert.ElementsMatch(t, []Node{"node1", "node2"}, nodes)
	})

	t.Run("should return all errors from write func", func(t *testing.T) {
		result := ReadAllFromChannel(Target.Execute(
			WithNodes([]Node{"", "", ""}),
			WithWriteFunc(WriteFuncWithResult([]error{errors.New("boom"), nil, errors.New("fail")})),
		))
		assert.ElementsMatch(t, []error{errors.New("boom"), nil, errors.New("fail")}, result.Value())
	})
}
