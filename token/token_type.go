package token

type TokenType struct {
	Name string
}

func (tokenType TokenType) GetName() string {
	return tokenType.Name
}

var (
	DelimiterType    = TokenType{"delimiter"}         // Разделитель
	IdentifierType   = TokenType{"identifier"}        // Идентификатор
	ConstantType     = TokenType{"const_number"}      // Шестнадцатиричное число
	AssignmentType   = TokenType{"assignment"}        // Присваивание
	AndType          = TokenType{"and"}               // and
	OrType           = TokenType{"or"}                // or
	XorType          = TokenType{"xor"}               // xor
	NotType          = TokenType{"not"}               // not
	LeftParenthType  = TokenType{"left_parentheses"}  // Скобки
	RightParenthType = TokenType{"right_parentheses"} // Скобки
	ErrorType        = TokenType{"error"}             // Ошибка
	CommentType      = TokenType{"comment"}           // Комментарий
	StartType        = TokenType{"start"}             // Начало
	EOFType          = TokenType{"EOF"}               // Конец
)
