package main

import (
	"goodhumored/lr1_object_code_generator/syntax_analyzer/precedence"
	"goodhumored/lr1_object_code_generator/token"
)

// Матрица предшествования
var precedenceMatrix = precedence.Matrix{
	token.IdentifierType:   precedence.Row{token.AssignmentType: precedence.Eq, token.RightParenthType: precedence.Gt, token.OrType: precedence.Gt, token.XorType: precedence.Gt, token.AndType: precedence.Gt, token.DelimiterType: precedence.Gt},
	token.AssignmentType:   precedence.Row{token.IdentifierType: precedence.Lt, token.LeftParenthType: precedence.Lt, token.RightParenthType: precedence.Lt, token.NotType: precedence.Lt, token.OrType: precedence.Lt, token.XorType: precedence.Lt, token.AndType: precedence.Lt, token.DelimiterType: precedence.Eq},
	token.LeftParenthType:  precedence.Row{token.IdentifierType: precedence.Lt, token.LeftParenthType: precedence.Lt, token.RightParenthType: precedence.Eq, token.NotType: precedence.Lt, token.OrType: precedence.Lt, token.XorType: precedence.Lt, token.AndType: precedence.Lt},
	token.RightParenthType: precedence.Row{token.RightParenthType: precedence.Gt, token.OrType: precedence.Gt, token.XorType: precedence.Gt, token.AndType: precedence.Gt, token.DelimiterType: precedence.Gt},
	token.NotType:          precedence.Row{token.LeftParenthType: precedence.Lt},
	token.OrType:           precedence.Row{token.IdentifierType: precedence.Lt, token.LeftParenthType: precedence.Lt, token.RightParenthType: precedence.Gt, token.NotType: precedence.Lt, token.OrType: precedence.Gt, token.XorType: precedence.Gt, token.AndType: precedence.Lt, token.DelimiterType: precedence.Gt},
	token.XorType:          precedence.Row{token.IdentifierType: precedence.Lt, token.LeftParenthType: precedence.Lt, token.RightParenthType: precedence.Gt, token.NotType: precedence.Lt, token.OrType: precedence.Gt, token.XorType: precedence.Gt, token.AndType: precedence.Lt, token.DelimiterType: precedence.Gt},
	token.AndType:          precedence.Row{token.IdentifierType: precedence.Lt, token.LeftParenthType: precedence.Lt, token.RightParenthType: precedence.Gt, token.NotType: precedence.Lt, token.OrType: precedence.Gt, token.XorType: precedence.Gt, token.AndType: precedence.Gt, token.DelimiterType: precedence.Gt},
	token.DelimiterType:    precedence.Row{token.IdentifierType: precedence.Gt},
}
