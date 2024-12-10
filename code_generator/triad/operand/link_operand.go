package operand

import (
	"fmt"
)

type LinkOperand struct{ LinkTo *Operand }

func (o LinkOperand) Hash() string {
	return fmt.Sprintf("^%v", o.LinkTo)
}

func (o LinkOperand) String() string {
	return fmt.Sprintf("^%v", o.LinkTo)
}

func (o LinkOperand) Value() (any, error) {
	return (*o.LinkTo).Value()
}

func Link(triad Operand) LinkOperand {
	return LinkOperand{
		LinkTo: &triad,
	}
}
