package rule

import (
	"fmt"

	"goodhumored/lr1_object_code_generator/syntax_analyzer/nonterminal"
)

// Интерфейс для представления символа
type Symbol interface {
	GetName() string
}

// Правило
type Rule struct {
	Left  nonterminal.NonTerminal
	Right [][]Symbol
}

// Метод получения строки из правила
func (r Rule) String() string {
	return fmt.Sprintf("%s -> %s", r.Left.GetName(), r.Right)
}
