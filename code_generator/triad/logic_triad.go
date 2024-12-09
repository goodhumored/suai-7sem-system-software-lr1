package triad

import "errors"

type LogicTriad struct {
	baseTriad
	operation func(left int, right int) int
}

func (t LogicTriad) Value() (any, error) {
	leftVal, err := t.left.Value()
	if err != nil {
		return 0, err
	}
	leftIntVal, ok := leftVal.(int)
	if !ok {
		return nil, errors.New("Bad value")
	}
	rightVal, err := t.right.Value()
	if err != nil {
		return 0, err
	}
	rightIntVal, ok := rightVal.(int)
	if !ok {
		return nil, errors.New("Bad value")
	}
	return t.operation(leftIntVal, rightIntVal), nil
}

func Logic(number int, left Operand, right Operand, operation func(left int, right int) int) LogicTriad {
	return LogicTriad{
		baseTriad: baseTriad{number: number, left: left, right: right},
		operation: operation,
	}
}
