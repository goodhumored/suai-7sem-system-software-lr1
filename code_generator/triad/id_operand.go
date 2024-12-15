package triad

type IdOperand struct{ name string }

func (o IdOperand) Hash() string {
	return o.name
}

func (o IdOperand) String() string {
	return o.name
}

func (o IdOperand) Value() (any, error) {
	return o.name, nil
}

func Id(name string) IdOperand {
	return IdOperand{
		name,
	}
}
