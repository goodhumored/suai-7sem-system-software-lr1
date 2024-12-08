package triad

type SameTriad[T any] struct {
	baseTriad[T]
}

func (t SameTriad[T]) Value() (T, error) {
	return t.left.Value()
}

func Same[T any](triad Triad[T], number int) SameTriad[T] {
	return SameTriad[T]{
		baseTriad[T]{number: number, left: triad, right: nil},
	}
}
