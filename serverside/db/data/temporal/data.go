package temporal

type hashableData interface {
	data
	hashable
}

type hashable interface {
	Hash() int
}

type data interface {
	Newness() int
}
