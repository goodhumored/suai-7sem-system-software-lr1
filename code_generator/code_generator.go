package codegenerator

import (
	"fmt"
	"strings"

	asm8086triadtranslator "goodhumored/lr1_object_code_generator/code_generator/asm_8086_triad_translator"
	code_optimisation "goodhumored/lr1_object_code_generator/code_generator/code_optimisation"
	"goodhumored/lr1_object_code_generator/syntax_analyzer/parse_tree"
)

func GenerateCode(tree parse_tree.ParseTree) (string, error) {
	var triadTranslator TriadTranslator
	triadTranslator = asm8086triadtranslator.Asm8086TriadTranslator{}

	triads := MapParseTreeToTriadList(tree)
	fmt.Println("Триады:")
	triads.Print()

	codeBeforeOptimisation, _ := triadTranslator.TranslateTriads(triads)
	code_optimisation.OptimiseCode(&triads)
	codeAfterOptimisation, err := triadTranslator.TranslateTriads(triads)

	printSavedLinesInfo(codeBeforeOptimisation, codeAfterOptimisation)

	return codeAfterOptimisation, err
}

func printSavedLinesInfo(codeBefore string, codeAfter string) {
	fmt.Printf("\nКод до оптимизации:\n%s\n", codeBefore)
	linesBeforeOptimisation := len(strings.Split(codeBefore, "\n"))
	linesAfterOptimisation := len(strings.Split(codeAfter, "\n"))
	difference := linesBeforeOptimisation - linesAfterOptimisation
	diffPerc := difference * 100 / linesBeforeOptimisation
	fmt.Printf("Количество оптимизированных операций (строк стало меньше на): %d (%d%%)\n\n", difference, diffPerc)
}
