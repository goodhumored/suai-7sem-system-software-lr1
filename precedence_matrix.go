package main

import (
	"goodhumored/lr1_object_code_generator/syntax_analyzer/precedence"
	"goodhumored/lr1_object_code_generator/token"
)

// Матрица предшествования
var precedenceMatrix = precedence.PrecedenceMatrix{
	token.IdentifierType:   map[token.TokenType]precedence.PrecedenceType{token.AssignmentType: precedence.Eq, token.RightParenthType: precedence.Gt, token.OrType: precedence.Gt, token.XorType: precedence.Gt, token.AndType: precedence.Gt, token.DelimiterType: precedence.Gt},
	token.AssignmentType:   map[token.TokenType]precedence.PrecedenceType{token.IdentifierType: precedence.Lt, token.LeftParenthType: precedence.Lt, token.RightParenthType: precedence.Lt, token.NotType: precedence.Lt, token.OrType: precedence.Lt, token.XorType: precedence.Lt, token.AndType: precedence.Lt, token.DelimiterType: precedence.Eq},
	token.LeftParenthType:  map[token.TokenType]precedence.PrecedenceType{token.IdentifierType: precedence.Lt, token.LeftParenthType: precedence.Lt, token.RightParenthType: precedence.Eq, token.NotType: precedence.Lt, token.OrType: precedence.Lt, token.XorType: precedence.Lt, token.AndType: precedence.Lt},
	token.RightParenthType: map[token.TokenType]precedence.PrecedenceType{token.RightParenthType: precedence.Gt, token.OrType: precedence.Gt, token.XorType: precedence.Gt, token.AndType: precedence.Gt, token.DelimiterType: precedence.Gt},
	token.NotType:          map[token.TokenType]precedence.PrecedenceType{token.LeftParenthType: precedence.Lt},
	token.OrType:           map[token.TokenType]precedence.PrecedenceType{token.IdentifierType: precedence.Lt, token.LeftParenthType: precedence.Lt, token.RightParenthType: precedence.Gt, token.NotType: precedence.Lt, token.OrType: precedence.Gt, token.XorType: precedence.Gt, token.AndType: precedence.Lt, token.DelimiterType: precedence.Gt},
	token.XorType:          map[token.TokenType]precedence.PrecedenceType{token.IdentifierType: precedence.Lt, token.LeftParenthType: precedence.Lt, token.RightParenthType: precedence.Gt, token.NotType: precedence.Lt, token.OrType: precedence.Gt, token.XorType: precedence.Gt, token.AndType: precedence.Lt, token.DelimiterType: precedence.Gt},
	token.AndType:          map[token.TokenType]precedence.PrecedenceType{token.IdentifierType: precedence.Lt, token.LeftParenthType: precedence.Lt, token.RightParenthType: precedence.Gt, token.NotType: precedence.Lt, token.OrType: precedence.Gt, token.XorType: precedence.Gt, token.AndType: precedence.Gt, token.DelimiterType: precedence.Gt},
}
