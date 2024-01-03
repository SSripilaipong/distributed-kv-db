package read

import "distributed-kv-db/common/typ"

func latestDataDummy[Data any]([]Data) Data {
	return typ.Zero[Data]()
}

func latestDataCaptureXs[Data any](xs *[]Data) func([]Data) Data {
	return func(ys []Data) Data {
		*xs = ys
		return latestDataDummy[Data](nil)
	}
}
