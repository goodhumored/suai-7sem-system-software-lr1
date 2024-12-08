package triad

import "errors"

type NotTriad struct {
	baseTriad[int]
}

func (t NotTriad) Value() (int, error) {
	if leftVal, err := t.left.Value(); err == nil {
		return ^leftVal, nil
	}
	return 0, errors.New("no value")
}

func Not(operand Operand[int], number int) NotTriad {
	return NotTriad{
		baseTriad[int]{
			left:   operand,
			number: number,
		},
	}
}
