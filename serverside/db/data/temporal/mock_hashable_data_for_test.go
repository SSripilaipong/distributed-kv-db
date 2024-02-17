package temporal

type hashableDataForTest struct {
	hash    int
	newness int
}

func (x hashableDataForTest) Hash() int {
	return x.hash
}

func (x hashableDataForTest) Newness() int {
	return x.newness
}

func newHashableDataForTest(newness int, hash int) DataWithHashComparison[hashableDataForTest] {
	return NewWithHashComparison(hashableDataForTest{
		hash:    hash,
		newness: newness,
	})
}
