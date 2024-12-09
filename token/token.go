package token

import "fmt"

// Структура Position представляет положение лексемы в строке
type Position struct {
	Start int
	End   int
}

// Структура Token представляет лексему с ее типом и значением
type Token struct {
	Type     TokenType // Тип
	value    string    // Значение
	Position Position  // Положение лексемы
}

// Функция получения имени токена, для соответствия интерфейсу символа
func (token Token) GetName() string {
	return token.Type.GetName()
}

func (token Token) Value() string {
	return token.value
}

// Фабричная функция для токенов, возвращающая замкнутую лямбда функцию для создания токена определённого типа
func tokenFactory(tokenType TokenType) func(string, Position) Token {
	return func(value string, position Position) Token {
		return Token{
			value:    value,
			Type:     tokenType,
			Position: position,
		}
	}
}

// Функция определяющая как токен преобразуется в строку
func (token Token) String() string {
	return fmt.Sprintf("%s (%s)", token.Type, token.Value())
}

// Функции создания лексем определённых типов
var (
	Delimiter    = tokenFactory(DelimiterType)           // Разделиитель
	Identifier   = tokenFactory(IdentifierType)          // Идентификатор
	Assignment   = tokenFactory(AssignmentType)          // Присваивание
	And          = tokenFactory(AndType)                 // И
	Or           = tokenFactory(OrType)                  // Или
	Xor          = tokenFactory(XorType)                 // Исключающее или
	Not          = tokenFactory(NotType)                 // Не
	LeftParenth  = tokenFactory(LeftParenthType)         // Левая скобка
	RightParenth = tokenFactory(RightParenthType)        // Правая скобка
	Error        = tokenFactory(ErrorType)               // Ошибка
	Comment      = tokenFactory(CommentType)             // Комментарий
	Start        = Token{StartType, "", Position{0, 0}}  // Начало строки
	EOF          = Token{EOFType, "EOF", Position{0, 0}} // Конец
)
