package triad

type LogicTriad struct {
	baseTriad[int]
	operation func(left int, right int) int
}

func (t LogicTriad) Value() (int, error) {
	leftVal, err := t.left.Value()
	if err != nil {
		return 0, err
	}
	rightVal, err := t.right.Value()
	if err != nil {
		return 0, err
	}
	return t.operation(leftVal, rightVal), nil
}

func Logic(number int, left Operand[int], right Operand[int], operation func(left int, right int) int) LogicTriad {
	return LogicTriad{
		baseTriad: baseTriad[int]{number: number, left: left, right: right},
		operation: operation,
	}
}
