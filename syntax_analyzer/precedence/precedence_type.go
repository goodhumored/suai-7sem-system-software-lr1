package precedence

// Тип предшествования
type PrecedenceType struct {
	Name string
}

// Типы предшествования
var (
	Lt        = PrecedenceType{"<"} // Предшествует
	Eq        = PrecedenceType{"="} // Составляет основу
	Gt        = PrecedenceType{">"} // Следует
	Undefined = PrecedenceType{"-"} // Неопределено
)
