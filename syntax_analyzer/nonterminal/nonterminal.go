package nonterminal

// Структура представляющая нетерминалы
type NonTerminal struct {
	Name string
}

// Метод для соответствия нетерменалов интерфейсу символ
func (nt NonTerminal) GetName() string {
	return nt.Name
}

func (nt NonTerminal) Value() string {
	return ""
}

var (
	E           = NonTerminal{"E"}                // Стандартный нетерминал
	Parenthesis = NonTerminal{"PARENTHESIS"}      // Стандартный нетерминал
	Assignment  = NonTerminal{"ASSIGNMENT"}       // Стандартный нетерминал
	Binary      = NonTerminal{"BINARY_OPERATION"} // Стандартный нетерминал
	Value       = NonTerminal{"VALUE"}            // Стандартный нетерминал
	Unary       = NonTerminal{"UNARY_OPERATION"}  // Стандартный нетерминал
	Root        = NonTerminal{"/"}                // Корневой нетерминал
)
