package readlatest

import (
	"context"
	"distributed-kv-db/common/rslt"
	peerRead "distributed-kv-db/serverside/db/coordinator/quorum/read"
)

func newFuncWithReadQuorumOfNodesData[Key, Data any](readQuorum func(context.Context, Key) rslt.Of[[]Data]) peerRead.Func[Key, Data] {
	return newFuncWithReadQuorumAndLatestData(readQuorum, latestDataDummy[Data])
}

func newFuncWithReadQuorumAndLatestData[Key, Data any](readQuorum func(context.Context, Key) rslt.Of[[]Data], latestData func([]Data) rslt.Of[Data]) peerRead.Func[Key, Data] {
	return composeReadLatest(readQuorum, latestData)
}

func newFuncWithLatestData[Key, Data any](latestData func([]Data) rslt.Of[Data]) peerRead.Func[Key, Data] {
	return composeReadLatest(readQuorumDummy[Key, Data], latestData)
}
