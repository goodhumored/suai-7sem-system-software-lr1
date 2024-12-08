package token_analyzer

import (
	"regexp"

	"goodhumored/lr1_object_code_generator/token"
)

// Вспомогатеьлная структура для установки соответствия шаблонов лексем
// с их фабричными функциями
type TokenPattern struct {
	Pattern *regexp.Regexp
	Type    func(string, token.Position) token.Token
}

// Массив соответствий шаблонов лексем
var tokenPatterns = []TokenPattern{
	{regex("or"), token.Or},
	{regex("xor"), token.Xor},
	{regex("and"), token.And},
	{regex("not"), token.Not},
	{regex("(0x|[0-9$])[0-9a-fA-F]+"), token.Identifier},
	{regex("[a-zA-Z][a-zA-Z0-9]+"), token.Identifier},
	{regex(":="), token.Assignment},
	{regex("#.*"), token.Comment},
	{regex("[(]"), token.LeftParenth},
	{regex("[)]"), token.RightParenth},
	{regex(";"), token.Delimiter},
}

// вспомогательная функция создающая объект регулярного выражения
// добавляющая в начале шаблона признак начала строки
func regex(pattern string) *regexp.Regexp {
	return regexp.MustCompile("^" + pattern)
}
