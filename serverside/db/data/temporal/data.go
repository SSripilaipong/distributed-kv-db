package temporal

type Hashable interface {
	data
	hashable
}

type hashable interface {
	Hash() int
}

type data interface {
	Newness() int
}
