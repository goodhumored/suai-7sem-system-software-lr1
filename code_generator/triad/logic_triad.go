package triad

import (
	"errors"
	"strconv"
)

type LogicTriad struct {
	baseTriad
	operation func(left int, right int) int
}

func (t LogicTriad) Value() (any, error) {
	leftIntVal, ok := parseOperand(t.left)
	if !ok {
		return nil, errors.New("failed parsing left")
	}
	rightIntVal, ok := parseOperand(t.right)
	if !ok {
		return nil, errors.New("failed parsing right")
	}
	value := t.operation(leftIntVal, rightIntVal)
	return strconv.Itoa(value), nil
}

func parseOperand(operand Operand) (int, bool) {
	if _, isId := operand.(IdOperand); !isId {
		return 0, false
	}
	val, err := operand.Value()
	if err != nil {
		return 0, false
	}
	strVal, ok := val.(string)
	if !ok {
		return 0, false
	}
	intVal, err := convertStrToNumber(strVal)
	if err != nil {
		return 0, false
	}
	return int(intVal), true
}

func convertStrToNumber(str string) (int64, error) {
	base := 10
	if len(str) > 1 && str[1] == 'x' {
		str = str[2:]
		base = 16
	}
	return strconv.ParseInt(str, base, 32)
}

func Logic(number int, left Operand, right Operand, operation func(left int, right int) int) LogicTriad {
	return LogicTriad{
		baseTriad: baseTriad{number: number, left: left, right: right},
		operation: operation,
	}
}
