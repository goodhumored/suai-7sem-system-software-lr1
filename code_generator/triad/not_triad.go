package triad

import (
	"errors"
	"fmt"

	"goodhumored/lr1_object_code_generator/code_generator/triad/operand"
)

type NotTriad struct {
	baseTriad
}

func (t NotTriad) String() string {
	return fmt.Sprintf("not(%s,)", t.left.String())
}

func (t NotTriad) Hash() string {
	return t.left.Hash()
}

func (t NotTriad) Value() (any, error) {
	if leftVal, err := t.left.Value(); err == nil {
		if val, ok := leftVal.(int); ok {
			return ^val, nil
		}
		return nil, errors.New("Bad value")
	}
	return 0, errors.New("no value")
}

func Not(operand operand.Operand, number int) NotTriad {
	return NotTriad{
		baseTriad{
			left:   operand,
			number: number,
		},
	}
}
