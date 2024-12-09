package triad

import "strconv"

type ConstantOperand struct{ value int }

func (o ConstantOperand) Value() (any, error) {
	return o.value, nil
}

func (o ConstantOperand) Hash() string {
	return strconv.Itoa(o.value)
}

func Constant(value int) ConstantOperand {
	return ConstantOperand{
		value,
	}
}
