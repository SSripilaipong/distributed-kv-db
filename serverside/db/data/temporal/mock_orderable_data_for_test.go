package temporal

type dataMock struct {
	newness int
	hash    int
}

func (d dataMock) Newness() int {
	return d.newness
}

func (d dataMock) Hash() int {
	return d.hash
}

func dataMockWithNewness(newness int) dataMock {
	return dataMock{newness: newness}
}

func dataMockWithNewnessAndHash(newness int, hash int) dataMock {
	return dataMock{newness: newness, hash: hash}
}
