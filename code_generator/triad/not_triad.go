package triad

import (
	"errors"
	"fmt"
	"strconv"
)

type NotTriad struct {
	baseTriad
}

func (t NotTriad) String() string {
	return fmt.Sprintf("not(%s,)", t.left.String())
}

func (t NotTriad) Hash() string {
	return fmt.Sprintf("not_%s", t.left.Hash())
}

func (t NotTriad) Value() (any, error) {
	if leftVal, err := t.left.Value(); err == nil {
		if strVal, ok := leftVal.(string); ok {
			strVal = strVal[2:]
			intVal, err := strconv.ParseInt(strVal, 16, 32)
			if err != nil {
				return nil, err
			}
			return strconv.Itoa(int(^intVal)), nil
		}
	}
	return 0, errors.New("no value")
}

func Not(operand Operand, number int) NotTriad {
	return NotTriad{
		baseTriad{
			left:   operand,
			number: number,
		},
	}
}
