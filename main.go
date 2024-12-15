package main

import (
	"fmt"
	"os"

	codegenerator "goodhumored/lr1_object_code_generator/code_generator"
	"goodhumored/lr1_object_code_generator/syntax_analyzer"
	"goodhumored/lr1_object_code_generator/token_analyzer"
)

func main() {
	source := getInput("./input-subexpr-test.txt") // читаем файл

	// выводим содержимое
	println("Содержимое входного файла:\n")
	fmt.Println(source)

	// запускаем распознание лексем
	tokenTable := token_analyzer.RecogniseTokens(source)

	// выводим лексемы
	fmt.Println("Таблица лексем:")
	fmt.Println(tokenTable)

	// Проверяем на ошибки
	if errors := tokenTable.GetErrors(); len(errors) > 0 {
		fmt.Printf("Во время лексического анализа было обнаружено: %d ошибок:\n", len(errors))
		for _, error := range errors {
			fmt.Printf("Неожиданный символ '%s'\n", error.Value())
		}
		return
	}

	// запускаем синтаксический анализатор
	tree, err := syntax_analyzer.AnalyzeSyntax(rulesTable, *tokenTable, precedenceMatrix)
	if err != nil {
		fmt.Printf("Ошибка при синтаксическом анализе строки: %s", err)
		return
	} else {
		fmt.Println("Строка принята!!!")
		tree.Print()
	}

	// запускаем генерацию объедкного кода
	code, err := codegenerator.GenerateCode(tree)
	if err != nil {
		fmt.Printf("Во время генерации кода возникли следующие ошибки: %v", err)
	}
	fmt.Printf("Результирующий код:\n%v\n", code)
	fmt.Printf("Исходный код: \n%s", source)
}

// Читает файл с входными данными, вызывает панику в случае неудачи
func getInput(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(data)
}
