package temporal

type hashableDataMock struct {
	newness int
	hash    int
}

func (d hashableDataMock) Newness() int {
	return d.newness
}

func (d hashableDataMock) Hash() int {
	return d.hash
}

func orderableDataWithNewness(newness int) hashableDataMock {
	return hashableDataMock{newness: newness}
}

func orderableDataWithNewnessAndHash(newness int, hash int) hashableDataMock {
	return hashableDataMock{newness: newness, hash: hash}
}
