package triad

import (
	"fmt"

	"goodhumored/lr1_object_code_generator/code_generator/triad/operand"
)

type OrTriad struct {
	LogicTriad
}

func (t OrTriad) String() string {
	return fmt.Sprintf("or(%s,%s)", t.left.String(), t.right.String())
}

func Or(left operand.Operand, right operand.Operand, number int) OrTriad {
	return OrTriad{
		LogicTriad: Logic(number, left, right, func(left int, right int) int {
			return left | right
		}),
	}
}
