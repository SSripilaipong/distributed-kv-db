package readlatest

type orderableDataMock struct {
	newness int
	hash    int
}

func (d orderableDataMock) Newness() int {
	return d.newness
}

func (d orderableDataMock) Hash() int {
	return d.hash
}

func orderableDataWithNewness(newness int) orderableDataMock {
	return orderableDataMock{newness: newness}
}

func orderableDataWithNewnessAndHash(newness int, hash int) orderableDataMock {
	return orderableDataMock{newness: newness, hash: hash}
}
