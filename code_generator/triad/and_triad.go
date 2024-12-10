package triad

import (
	"fmt"

	"goodhumored/lr1_object_code_generator/code_generator/triad/operand"
)

type AndTriad struct {
	LogicTriad
}

func (t AndTriad) String() string {
	return fmt.Sprintf("and(%s,%s)", t.left.String(), t.right.String())
}

func And(left operand.Operand, right operand.Operand, number int) AndTriad {
	return AndTriad{
		LogicTriad: Logic(number, left, right, func(left int, right int) int {
			return left & right
		}),
	}
}
