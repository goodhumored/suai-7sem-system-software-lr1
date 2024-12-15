package triad

import (
	"fmt"
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

func Assignment(left Operand, right Operand, number int) AssignmentTriad {
	return AssignmentTriad{
		baseTriad{left: left, right: right, number: number},
	}
}
