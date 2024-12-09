package codegenerator

import (
	"fmt"

	"goodhumored/lr1_object_code_generator/syntax_analyzer/parse_tree"
)

func GenerateCode(tree parse_tree.ParseTree) string {
	triads := MapParseTreeToTriadList(tree)
	for i, triad := range triads {
		fmt.Printf("%d)%s\n", i, triad.String())
	}
	return ""
}
