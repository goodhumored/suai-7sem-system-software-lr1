package syntax_analyzer

import (
	"errors"
	"fmt"

	"goodhumored/lr1_object_code_generator/syntax_analyzer/nonterminal"
	"goodhumored/lr1_object_code_generator/syntax_analyzer/parse_tree"
	"goodhumored/lr1_object_code_generator/syntax_analyzer/precedence"
	"goodhumored/lr1_object_code_generator/syntax_analyzer/rule"
	"goodhumored/lr1_object_code_generator/token"
	"goodhumored/lr1_object_code_generator/token_table"
)

// Функция для анализа синтаксиса, принимает таблицу токенов, список правил и матрицу предшествования
func AnalyzeSyntax(ruleTable rule.RuleTable, tokenTable token_table.TokenTable, matrix precedence.Matrix) (parse_tree.ParseTree, error) {
	// Создаём дерево
	rootNode := parse_tree.CreateNode(nonterminal.Root)
	tree := parse_tree.ParseTree{Root: &rootNode}
	// Получаем лексемы из таблицы
	tokens := tokenTable.GetTokens()
	tokenIndex := 1
	// Создаём стек
	stack := symbolStack{tokens[0]}

	for {
		// Берём ближайший к вершине терминал
		stackTerminal := stack.PeekTopTerminal()
		// Берём текущий символ входной строки
		if len(tokens) <= tokenIndex {
			return tree, errors.New("Токены закончились, до конца свернуть не удалось")
		}
		inputToken := tokens[tokenIndex]
		// Если строка принята, значит возвращаем дерево вывода
		if isInputAccepted(inputToken, stack) {
			return tree, nil
		}
		// Если комментарий - пропускаем
		if inputToken.Type == token.CommentType {
			tokenIndex += 1
			continue
		}

		fmt.Printf("Лексема: '%s' \n", tokens[tokenIndex].Value())
		fmt.Printf("Стек: %s \n", stack)

		// Получаем предшествование из матрицы
		prec := matrix.GetPrecedence(stackTerminal.Type, inputToken.Type)

		// Если предшествование или =, тогда сдвигаем
		if prec == precedence.Lt || prec == precedence.Eq {
			print("Сдвигаем\n")
			node := &parse_tree.Node{Value: inputToken.Value(), Symbol: inputToken, Children: []*parse_tree.Node{}}
			tree.AddNode(node) // Добавляем узел в дерево
			stack = stack.Push(inputToken)
			tokenIndex += 1
		} else if prec == precedence.Gt { // Иначе сворачиваем
			print("Сворачиваем\n")
			// сворачиваем стек
			newStack, rule, err := reduce(stack, ruleTable)
			if err != nil {
				return tree, err
			}
			stack = newStack
			// сворачиваем дерево
			tree.Reduce(*rule)
		} else {
			// Если предшествование не определено - выдаем ошибку
			return tree, fmt.Errorf("Ошибка в синтексе, неожиданное сочетание символов %s и %s (%d)", stackTerminal.GetName(), inputToken.GetName(), inputToken.Position.End)
		}
		println("==============")
	}
}

// Проверка на завершённость
func isInputAccepted(currentToken token.Token, stack symbolStack) bool {
	nextTerminal := stack.PeekTopTerminal()
	nextSymbol := stack.Peek()
	return currentToken.Type == token.EOFType && // Если дошли до конца строки
		nextTerminal != nil &&
		nextTerminal.Type == token.Start.Type && // Если ближайший терминал в стеке - начало строки
		nextSymbol != nil &&
		nextSymbol == nonterminal.Assignment // А на вершине строки - целевой символ
}

// Функция свёртки стека
func reduce(stack symbolStack, ruleTable rule.RuleTable) (symbolStack, *rule.Rule, error) {
	for {
		// Если есть применимое к стеку правило
		if rule := ruleTable.GetRuleByRightSide(stack); rule != nil {
			fmt.Printf("Нашлось правило: %v, пушим %s в стек\n", rule, rule.Left)
			// обновляем стек
			stack = append(stack[:len(stack)-len(rule.Right)], rule.Left)
			return stack, rule, nil
		} else {
			// Если нет выдаем ошибку
			return stack, nil, fmt.Errorf("Не найдено правил для свёртки")
		}
	}
}
