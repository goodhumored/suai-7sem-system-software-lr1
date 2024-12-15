package precedence

import (
	"goodhumored/lr1_object_code_generator/token"
)

type Row = map[token.TokenType]PrecedenceType

// Матрица предшествования
type Matrix map[token.TokenType]Row

// Метод для поиска типа предшествования для двух терминалов
func (matrix Matrix) GetPrecedence(left, right token.TokenType) PrecedenceType {
	// Если левый символ - начало файла, возвращаем предшествие
	if left == token.StartType {
		return Lt
	}
	// Если правый символ - конец файла, возвращаем следствие
	if right == token.EOFType {
		return Gt
	}
	// Если находится - возвращаем
	if val, ok := matrix[left]; ok {
		if precedence, ok := val[right]; ok {
			return precedence
		}
	}
	// Если не находится - возвращаем неопределённость
	return Undefined
}
