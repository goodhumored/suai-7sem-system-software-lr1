package triad

import (
	"fmt"
)

type OrTriad struct {
	LogicTriad
}

func (t OrTriad) Hash() string {
	return fmt.Sprintf("or_%s_%s", t.left.Hash(), t.right.Hash())
}

func (t OrTriad) String() string {
	return fmt.Sprintf("or(%s,%s)", t.left.String(), t.right.String())
}

func Or(left Operand, right Operand, number int) OrTriad {
	return OrTriad{
		LogicTriad: Logic(number, left, right, func(left int, right int) int {
			return left | right
		}),
	}
}
