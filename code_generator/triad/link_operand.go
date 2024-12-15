package triad

import (
	"fmt"
)

type LinkOperand struct{ LinkTo int }

func (o LinkOperand) Hash() string {
	return fmt.Sprintf("^%d", o.LinkTo)
}

func (o LinkOperand) String() string {
	return fmt.Sprintf("^%d", o.LinkTo)
}

func (o LinkOperand) Value() (any, error) {
	return o.LinkTo, nil
}

func Link(triad Triad) LinkOperand {
	return LinkOperand{
		LinkTo: triad.Number(),
	}
}
