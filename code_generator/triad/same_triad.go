package triad

type SameTriad struct {
	baseTriad
}

func (t SameTriad) Value() (any, error) {
	return t.left.Value()
}

func (t SameTriad) Hash() string {
	return t.left.Hash()
}

func Same(triad Triad, number int) SameTriad {
	return SameTriad{
		baseTriad{number: number, left: triad, right: nil},
	}
}
