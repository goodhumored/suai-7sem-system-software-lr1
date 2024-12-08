package triad

type ConstantTriad[T any] struct {
	baseTriad[T]
	value T
}

func (t ConstantTriad[T]) Value() (T, error) {
	return t.value, nil
}

func C[T any](number int, value T) ConstantTriad[T] {
	return ConstantTriad[T]{
		baseTriad: baseTriad[T]{number: number, left: nil, right: nil},
		value:     value,
	}
}
