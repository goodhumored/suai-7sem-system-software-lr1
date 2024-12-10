package triad

import (
	"fmt"

	"goodhumored/lr1_object_code_generator/code_generator/triad/operand"
)

type AssignmentTriad struct {
	baseTriad
}

func (t AssignmentTriad) String() string {
	return fmt.Sprintf(":=(%s,%s)", t.left.String(), t.right.String())
}

func (t AssignmentTriad) Value() (any, error) {
	return t.right.Value()
}

func Assignment(left operand.Operand, right operand.Operand, number int) AssignmentTriad {
	return AssignmentTriad{
		baseTriad{left: left, right: right, number: number},
	}
}
