package syntax_analyzer

import (
	"fmt"

	"goodhumored/lr1_object_code_generator/syntax_analyzer/rule"
	"goodhumored/lr1_object_code_generator/token"
)

// Стек символов
type symbolStack []rule.Symbol

// Добавление символа в стек
func (s symbolStack) Push(e rule.Symbol) symbolStack {
	return append(s, e)
}

// Просмотр верхнего элемента стека
func (s symbolStack) Peek() rule.Symbol {
	length := len(s)
	if length == 0 {
		return nil
	}
	return s[length-1]
}

// Просмотр n-ного элемента стека
func (s symbolStack) PeekN(n int) rule.Symbol {
	length := len(s)
	if length == 0 {
		return nil
	}
	return s[length-n-1]
}

// Поиск ближайшего к вершине терминала в стеке
func (s symbolStack) PeekTopTerminal() *token.Token {
	for i := range s {
		symbol := s.PeekN(i)
		if token, ok := symbol.(token.Token); ok {
			return &token
		}
	}
	return nil
}

// Вспомогательный метод преобразования стека символов в строку
func (s symbolStack) String() string {
	str := ""
	for _, i := range s {
		str += fmt.Sprintf("%s ", i.GetName())
	}
	return str
}

// Вспомогательный метод вывода стека символов
func (s symbolStack) Print() {
	fmt.Print(s.String())
}
