package main

import (
	"goodhumored/lr1_object_code_generator/syntax_analyzer/nonterminal"
	"goodhumored/lr1_object_code_generator/syntax_analyzer/rule"
	"goodhumored/lr1_object_code_generator/token"
)

func or(symbols ...rule.Symbol) []rule.Symbol {
	return symbols
}

var (
	valueSymbols           = or(nonterminal.Binary, nonterminal.Unary, token.IdentifierType, token.ConstantType, nonterminal.Parenthesis, nonterminal.Value)
	binaryOperatorsSymbols = or(token.AndType, token.OrType, token.XorType)
)

// Правила грамматики
var assignmentRule = rule.Rule{
	Left:  nonterminal.Assignment,
	Right: [][]rule.Symbol{or(token.IdentifierType), or(token.AssignmentType), valueSymbols, or(token.DelimiterType)},
}

var binaryRule = rule.Rule{
	Left:  nonterminal.Binary,
	Right: [][]rule.Symbol{valueSymbols, binaryOperatorsSymbols, valueSymbols},
}

var unaryRule = rule.Rule{
	Left:  nonterminal.Unary,
	Right: [][]rule.Symbol{or(token.NotType), or(token.LeftParenthType), valueSymbols, or(token.RightParenthType)},
}

var parenthesisRule = rule.Rule{
	Left:  nonterminal.Parenthesis,
	Right: [][]rule.Symbol{or(token.LeftParenthType), valueSymbols, or(token.RightParenthType)},
}

var identifierRule = rule.Rule{
	Left:  nonterminal.Value,
	Right: [][]rule.Symbol{valueSymbols},
}

var rootRule = rule.Rule{
	Left:  nonterminal.Root,
	Right: [][]rule.Symbol{or(token.StartType), or(nonterminal.Assignment), or(token.EOFType)},
}

var rulesTable = rule.RuleTable{Rules: []rule.Rule{
	unaryRule,
	parenthesisRule,
	binaryRule,
	assignmentRule,
	rootRule,
	identifierRule,
}}
