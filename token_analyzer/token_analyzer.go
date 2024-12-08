package token_analyzer

import (
	"strings"

	"goodhumored/lr1_object_code_generator/token"
	"goodhumored/lr1_object_code_generator/token_table"
)

// Распознаёт токены в данной строке построчно и записывает в таблицу
func RecogniseTokens(source string) *token_table.TokenTable {
	tokenTable := &token_table.TokenTable{}
	tokenTable.Add(token.Start)
	for _, line := range strings.Split(source, "\n") {
		recogniseTokensLine(line, tokenTable)
	}
	tokenTable.Add(token.EOF)
	return tokenTable
}

// Распознаёт лексемы в данной строке и записывает в таблицу
func recogniseTokensLine(line string, tokenTable *token_table.TokenTable) {
	for {
		line = strings.Trim(line, " ") // обрезаем пробельные символы в строке
		if len(line) == 0 {            // если строка пустая - завершаем обработку строки
			return
		}
		nextToken := getNextToken(line)      // ищем очередную лексему
		tokenTable.Add(nextToken)            // добавляем лексему в таблицу
		line = line[nextToken.Position.End:] // вырезаем обработанную часть
	}
}

// Ищет очередную лексему в строке
func getNextToken(str string) token.Token {
	// проходим по всем шаблонам лексем
	for _, tokenPattern := range tokenPatterns {
		res := tokenPattern.Pattern.FindStringIndex(str)
		if res != nil {
			return tokenPattern.Type(str[res[0]:res[1]], token.Position{res[0], res[1]})
		}
	}
	return token.Error(str[0:1], token.Position{0, 1})
}
