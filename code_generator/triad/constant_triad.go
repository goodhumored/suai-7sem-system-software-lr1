package triad

import "fmt"

type ConstantTriad struct {
	baseTriad
	value any
}

func (t ConstantTriad) Value() (any, error) {
	return t.value, nil
}

func (t ConstantTriad) String() string {
	return fmt.Sprintf("C(%s,)", t.value.(string))
}

func C[T any](number int, value T) ConstantTriad {
	return ConstantTriad{
		baseTriad: baseTriad{number: number, left: nil, right: nil},
		value:     value,
	}
}
