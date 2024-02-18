package readlatest

import (
	"context"
	"distributed-kv-db/common/fn"
	"distributed-kv-db/common/rslt"
	quorumRead "distributed-kv-db/serverside/db/coordinator/quorum/read"
	temporalData "distributed-kv-db/serverside/db/data/temporal"
)

func New[Key any, Data temporalData.Hashable](readQuorum func(context.Context, Key) rslt.Of[[]Data]) quorumRead.Func[Key, Data] {
	return composeReadLatest[Key, Data](readQuorum, temporalData.LatestInSlice[Data])
}

func composeReadLatest[Key, Data any](
	readQuorum func(context.Context, Key) rslt.Of[[]Data],
	latestData func([]Data) rslt.Of[Data],
) quorumRead.Func[Key, Data] {
	return fn.Uncurry(func(ctx context.Context) func(key Key) rslt.Of[Data] {
		return fn.Compose(
			rslt.FmapPartial(latestData), fn.WithArg(ctx, readQuorum),
		)
	})
}
