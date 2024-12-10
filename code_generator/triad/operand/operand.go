package operand

type Operand interface {
	Value() (any, error)
	Hash() string
	String() string
}
