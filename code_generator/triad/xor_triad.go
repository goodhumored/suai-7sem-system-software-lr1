package triad

import (
	"fmt"

	"goodhumored/lr1_object_code_generator/code_generator/triad/operand"
)

type XorTriad struct {
	LogicTriad
}

func (t XorTriad) String() string {
	return fmt.Sprintf("xor(%s,%s)", t.left.String(), t.right.String())
}

func Xor(left operand.Operand, right operand.Operand, number int) XorTriad {
	return XorTriad{
		LogicTriad: Logic(number, left, right, func(left int, right int) int {
			return left ^ right
		}),
	}
}
