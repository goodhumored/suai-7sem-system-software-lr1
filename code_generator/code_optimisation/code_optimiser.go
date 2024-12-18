package code_optimisation

import (
	"fmt"

	"goodhumored/lr1_object_code_generator/code_generator/triad"
)

func OptimiseCode(triadList *triad.TriadList) {
	foldConstants(triadList)
	fmt.Println("\nПосле свёртки констант:")
	triadList.Print()
	println()
	eliminateCommonSubexpression(triadList)
	fmt.Println("После удаления лишнего кода:")
	triadList.Print()
	println()
}
