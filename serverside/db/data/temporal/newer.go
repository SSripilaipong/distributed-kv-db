package temporal

func Newer[D Data](x, y D) D {
	xNewness, yNewness := x.Newness(), y.Newness()

	switch {
	case xNewness > yNewness:
		return x
	case xNewness < yNewness:
		return y
	case x.Hash() > y.Hash():
		return x
	default:
		return y
	}
}
