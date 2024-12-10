package operand

import (
	"fmt"
)

type LinkOperand struct{ linkTo *Operand }

func (o LinkOperand) Hash() string {
	return (*o.linkTo).Hash()
}

func (o LinkOperand) String() string {
	return fmt.Sprintf("^%v", o.linkTo)
}

func (o LinkOperand) Value() (any, error) {
	return (*o.linkTo).Value()
}

func Link(triad Operand) LinkOperand {
	return LinkOperand{
		linkTo: &triad,
	}
}
