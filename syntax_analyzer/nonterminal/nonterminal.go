package nonterminal

// Структура представляющая нетерминалы
type NonTerminal struct {
	Name string
}

// Метод для соответствия нетерменалов интерфейсу символ
func (nt NonTerminal) GetName() string {
	return nt.Name
}

var (
	E    = NonTerminal{"E"} // Стандартный нетерминал
	Null = NonTerminal{"/"} // Корневой нетерминал
)
