package token_table

import (
	"fmt"
	"strings"

	"goodhumored/lr1_object_code_generator/token"
)

// Таблица лексем
type TokenTable struct {
	tokens []token.Token
}

// Метод добавления лексемы в таблицу
func (tt *TokenTable) Add(token token.Token) {
	tt.tokens = append(tt.tokens, token)
}

// Метод получения списка найденных лексем
func (tt TokenTable) GetTokens() []token.Token {
	return tt.tokens
}

// Вспомогательная функция для вывода таблицы
func (tt *TokenTable) Print() {
	errors := tt.GetErrors()
	if len(errors) > 0 {
		errorMsg := ""
		for _, error := range errors {
			errorMsg += fmt.Sprintf("Неизвестный символ: %s \n", error.Value)
		}
		fmt.Println(fmt.Errorf(errorMsg))
	}
	fmt.Println(tt.String())
}

// Вспомогательная функция для генерации строки с таблицей лексем
func (tt *TokenTable) String() string {
	if len(tt.tokens) == 0 {
		return "Ни одного токена не найдено"
	}

	// Определяем максимальную ширину столбца
	maxTypeLen := len("Тип")
	maxValueLen := len("Значение")
	for _, token := range tt.tokens {
		if len(token.Type.Name) > maxTypeLen {
			maxTypeLen = len(token.Type.Name)
		}
		if len(token.Value) > maxValueLen {
			maxValueLen = len(token.Value)
		}
	}

	// создаем шапку и рамки
	header := fmt.Sprintf("| %-*s | %-*s |", maxTypeLen, "Тип", maxValueLen, "Значение")
	border := fmt.Sprintf("+-%s-+-%s-+", strings.Repeat("-", maxTypeLen), strings.Repeat("-", maxValueLen))

	// Собираем таблицу
	res := border + "\n" + header + "\n" + border + "\n"
	for _, token := range tt.tokens {
		res += fmt.Sprintf("| %-*s | %-*s |\n", maxTypeLen, token.GetName(), maxValueLen, token.Value)
	}
	res += border

	return res
}

// Функция возвращающая все ошибки в таблице
func (tt TokenTable) GetErrors() []token.Token {
	tokens := []token.Token{}
	for _, recognisedToken := range tt.tokens {
		if recognisedToken.Type == token.ErrorType {
			tokens = append(tokens, recognisedToken)
		}
	}
	return tokens
}
