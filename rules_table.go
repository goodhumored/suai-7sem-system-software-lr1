package main

import (
	"goodhumored/lr1_object_code_generator/syntax_analyzer/nonterminal"
	"goodhumored/lr1_object_code_generator/syntax_analyzer/rule"
	"goodhumored/lr1_object_code_generator/token"
)

// Правила грамматики
var rulesTable = rule.RuleTable{Rules: []rule.Rule{
	{Left: nonterminal.E, Right: []rule.Symbol{token.IdentifierType, token.AssignmentType, nonterminal.E, token.DelimiterType}},
	{Left: nonterminal.E, Right: []rule.Symbol{nonterminal.E, token.OrType, nonterminal.E}},
	{Left: nonterminal.E, Right: []rule.Symbol{nonterminal.E, token.XorType, nonterminal.E}},
	{Left: nonterminal.E, Right: []rule.Symbol{nonterminal.E, token.AndType, nonterminal.E}},
	{Left: nonterminal.E, Right: []rule.Symbol{token.NotType, token.LeftParenthType, nonterminal.E, token.RightParenthType}},
	{Left: nonterminal.E, Right: []rule.Symbol{token.LeftParenthType, nonterminal.E, token.RightParenthType}},
	{Left: nonterminal.E, Right: []rule.Symbol{token.IdentifierType}},
}}
