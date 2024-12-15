package triad

import (
	"fmt"
)

type AndTriad struct {
	LogicTriad
}

func (t AndTriad) Hash() string {
	return fmt.Sprintf("and_%s_%s", t.left.Hash(), t.right.Hash())
}

func (t AndTriad) String() string {
	return fmt.Sprintf("and(%s,%s)", t.left.String(), t.right.String())
}

func And(left Operand, right Operand, number int) AndTriad {
	return AndTriad{
		LogicTriad: Logic(number, left, right, func(left int, right int) int {
			return left & right
		}),
	}
}
