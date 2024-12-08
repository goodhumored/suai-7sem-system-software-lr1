package triad

type AssignmentTriad[T any] struct {
	baseTriad[T]
}

func (t AssignmentTriad[T]) Value() (T, error) {
	return t.right.Value()
}

func Assignment[T any](left Operand[T], right Triad[T], number int) AssignmentTriad[T] {
	return AssignmentTriad[T]{
		baseTriad[T]{left: left, right: right, number: number},
	}
}
