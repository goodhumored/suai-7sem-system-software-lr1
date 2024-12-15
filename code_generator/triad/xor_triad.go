package triad

import (
	"fmt"
)

type XorTriad struct {
	LogicTriad
}

func (t XorTriad) String() string {
	return fmt.Sprintf("xor(%s,%s)", t.left.String(), t.right.String())
}

func Xor(left Operand, right Operand, number int) XorTriad {
	return XorTriad{
		LogicTriad: Logic(number, left, right, func(left int, right int) int {
			return left ^ right
		}),
	}
}
