package parse_tree

import (
	"fmt"

	"goodhumored/lr1_object_code_generator/syntax_analyzer/rule"
)

// Узел дерева вывода
type Node struct {
	Symbol   rule.Symbol
	Children []*Node
}

// Вспомогательная функция для создания пустого узла
func CreateNode(s rule.Symbol) Node {
	return Node{s, []*Node{}}
}

// Метод, добавляющий дочерний узел
func (n *Node) AddChild(child *Node) {
	n.Children = append(n.Children, child)
}

// Метод свёртки узла дерева
func (node *Node) Reduce(rule rule.Rule) bool {
	// Если не можем применить правило к текущему узлу - уходим
	if !node.CanApplyRule(rule) {
		return false
	}
	// считаем разницу длин правой части правила и детей узла
	lenDiff := len(node.Children) - len(rule.Right)

	// копируем слайс с нужными нам узлами, которые собираемся заменять
	nodes := make([]*Node, len(rule.Right))
	copy(nodes, node.Children[lenDiff:])

	// перезаписываем дочерние узлы узла
	node.Children = append(node.Children[:lenDiff], &Node{rule.Left, nodes})
	return true
}

// Функция проверки возможности применения правила к дочерним узлам узла
func (node Node) CanApplyRule(rule rule.Rule) bool {
	lenDiff := len(node.Children) - len(rule.Right)
	// Если в правиле больше элементов чем в узле - уходим
	if lenDiff < 0 {
		return false
	}
	// Проходимся по символам правила и сравниваем с дочерними символами
	for i, rule := range rule.Right {
		if rule.GetName() != node.Children[i+lenDiff].Symbol.GetName() {
			return false
		}
	}
	return true
}

// Метод для рекурсивного вывода узлов дерева в консоль
func (node *Node) Print(prefix string, isTail bool) {
	// Выводим символ узла с отступом
	var branch, prefixSuffix string
	if isTail {
		prefixSuffix = "    "
		branch = "└── "
	} else {
		branch = "├── "
		prefixSuffix = "│   "
	}
	fmt.Println(prefix + branch + node.Symbol.GetName())

	// Рекурсивно выводим дочерние узлы
	for i := 0; i < len(node.Children)-1; i++ {
		node.Children[i].Print(prefix+prefixSuffix, false)
	}
	if len(node.Children) > 0 {
		node.Children[len(node.Children)-1].Print(prefix+prefixSuffix, true)
	}
}
