package triad

type Triad interface {
	Operand
	Number() int
	SetNumber(nubmer int)
	Left() Operand
	SetLeft(Operand)
	SetRight(Operand)
	Right() Operand
	Hash() string
}

type baseTriad struct {
	number int
	left   Operand
	right  Operand
}

func (t baseTriad) Number() int {
	return t.number
}

func (t *baseTriad) SetNumber(number int) {
	t.number = number
}

func (t *baseTriad) SetLeft(newLeft Operand) {
	t.left = newLeft
}

func (t *baseTriad) SetRight(newRight Operand) {
	t.right = newRight
}

func (t baseTriad) Left() Operand {
	return t.left
}

func (t baseTriad) Right() Operand {
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
