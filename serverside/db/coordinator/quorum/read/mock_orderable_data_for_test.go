package read

type orderableDataMock struct {
	newness int
}

func (d orderableDataMock) Newness() int {
	return d.newness
}

func (d orderableDataMock) Hash() string {
	//TODO implement me
	panic("implement me")
}

func orderableDataWithNewness(newness int) orderableDataMock {
	return orderableDataMock{newness: newness}
}
