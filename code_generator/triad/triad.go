package triad

type Operand[T any] interface {
	Value() (T, error)
	Hash() string
}

type Triad[T any] interface {
	Operand[T]
	Number() int
	SetNumber(nubmer int)
	Left() Operand[T]
	Right() Operand[T]
	Hash() string
}

type baseTriad[T any] struct {
	number int
	left   Operand[T]
	right  Operand[T]
}

func (t baseTriad[T]) Number() int {
	return t.number
}

func (t *baseTriad[T]) SetNumber(number int) {
	t.number = number
}

func (t baseTriad[T]) Left() Operand[T] {
	return t.left
}

func (t baseTriad[T]) Right() Operand[T] {
	return t.right
}

func (t baseTriad[T]) Hash() string {
	return t.left.Hash() + "_" + t.right.Hash()
}
