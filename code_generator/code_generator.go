package codegenerator

import (
	"fmt"

	asm8086triadtranslator "goodhumored/lr1_object_code_generator/code_generator/asm_8086_triad_translator"
	"goodhumored/lr1_object_code_generator/syntax_analyzer/parse_tree"
)

func GenerateCode(tree parse_tree.ParseTree) (string, error) {
	triads := MapParseTreeToTriadList(tree)
	for _, triad := range triads.Triads() {
		fmt.Printf("%d)%s\n", triad.Number(), triad.String())
	}
	var triadTranslator TriadTranslator
	triadTranslator = asm8086triadtranslator.Asm8086TriadTranslator{}
	return triadTranslator.TranslateTriads(triads)
}
