package codegenerator

import (
	asm8086triadtranslator "goodhumored/lr1_object_code_generator/code_generator/asm_8086_triad_translator"
	code_optimisation "goodhumored/lr1_object_code_generator/code_generator/code_optimisation"
	"goodhumored/lr1_object_code_generator/syntax_analyzer/parse_tree"
)

func GenerateCode(tree parse_tree.ParseTree) (string, error) {
	triads := MapParseTreeToTriadList(tree)
	triads.Print()
	code_optimisation.OptimiseCode(&triads)
	var triadTranslator TriadTranslator
	triadTranslator = asm8086triadtranslator.Asm8086TriadTranslator{}
	return triadTranslator.TranslateTriads(triads)
}
