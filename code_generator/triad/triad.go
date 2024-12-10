package triad

import "goodhumored/lr1_object_code_generator/code_generator/triad/operand"

type Triad interface {
	operand.Operand
	Number() int
	SetNumber(nubmer int)
	Left() operand.Operand
	Right() operand.Operand
	Hash() string
}

type baseTriad struct {
	number int
	left   operand.Operand
	right  operand.Operand
}

func (t baseTriad) Number() int {
	return t.number
}

func (t *baseTriad) SetNumber(number int) {
	t.number = number
}

func (t baseTriad) Left() operand.Operand {
	return t.left
}

func (t baseTriad) Right() operand.Operand {
	return t.right
}

func (t baseTriad) Hash() string {
	hash := ""
	if t.left != nil {
		hash += t.left.Hash()
	}
	if t.right != nil {
		hash += "_" + t.right.Hash()
	}
	return hash
}
