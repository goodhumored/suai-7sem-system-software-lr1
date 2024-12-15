package triad

import "fmt"

type SameTriad struct {
	baseTriad
	SameAs Triad
}

func (t SameTriad) Value() (any, error) {
	return t.left.Value()
}

func (t SameTriad) Hash() string {
	return t.left.Hash()
}

func (t SameTriad) String() string {
	return fmt.Sprintf("Same(%d,)", t.SameAs.Number())
}

func Same(triad Triad, number int) SameTriad {
	return SameTriad{
		baseTriad: baseTriad{number: number, left: triad, right: nil},
		SameAs:    triad,
	}
}
