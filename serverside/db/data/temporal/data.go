package temporal

type Data interface {
	Newness() int
	Hash() int
}
