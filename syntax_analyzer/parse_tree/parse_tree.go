package parse_tree

import (
	"fmt"

	"goodhumored/lr1_object_code_generator/syntax_analyzer/rule"
)

// Дерево вывода
type ParseTree struct {
	Root *Node
}

// Метод добавления узлов в дерево
func (tree *ParseTree) AddNode(node *Node) {
	tree.Root.AddChild(node)
}

// Метод для свёртки дерева по правилу
func (tree *ParseTree) Reduce(rule rule.Rule) {
	fmt.Printf("Применяем правило %s к дереву\n", rule)
	if tree.Root.Reduce(rule) {
		fmt.Printf("Успешно применено\n")
	} else {
		fmt.Printf("Правило %s применить не удалось\n", rule)
	}
}

// Метод для вывода дерева
func (tree ParseTree) Print() {
	tree.Root.Print("", true)
}
