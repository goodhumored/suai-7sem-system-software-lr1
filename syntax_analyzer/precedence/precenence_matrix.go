package precedence

import (
	"fmt"

	"goodhumored/lr1_object_code_generator/token"
)

// Матрица предшествования
type PrecedenceMatrix map[token.TokenType]map[token.TokenType]PrecedenceType

// Метод для поиска типа предшествования для двух терминалов
func (matrix PrecedenceMatrix) GetPrecedence(left, right token.TokenType) PrecedenceType {
	fmt.Printf("rigght: %v, left: %v", right, left)
	// Если левый символ - начало файла, возвращаем предшествие
	if left == token.StartType {
		return Lt
	}
	// Если правый символ - конец файла, возвращаем следствие
	if right == token.EOFType {
		fmt.Print("aaaaaaaaaaaaaa")
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
