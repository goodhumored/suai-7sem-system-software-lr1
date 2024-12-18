package code_optimisation

import (
	"goodhumored/lr1_object_code_generator/code_generator/triad"
)

type (
	dependencyTable          = map[string]int
	triadDependencyTableItem struct {
		triad            triad.Triad
		dependencyNumber int
	}
)

func eliminateCommonSubexpression(triads *triad.TriadList) {
	operandTable := dependencyTable{}
	triadTable := map[string]triadDependencyTableItem{}
	for i, t := range triads.Triads() {
		updateDep := true
		// fmt.Printf("\ntriad: %v\n", t)
		leftOperand, ok := checkSameOperand(t.Left(), *triads)
		if ok {
			t.SetLeft(leftOperand)
		}
		rightOperand, ok := checkSameOperand(t.Right(), *triads)
		if ok {
			t.SetRight(rightOperand)
		}
		leftDependency := getOperandDep(leftOperand, operandTable, *triads)
		rightDependency := getOperandDep(rightOperand, operandTable, *triads)

		triadDependency := max(leftDependency, rightDependency) + 1
		// fmt.Printf("hash is %s\n", t.Hash())
		if dependency, ok := triadTable[t.Hash()]; ok {
			// fmt.Printf("have dep: %v, calculated is %v\n", dependency, triadDependency)
			if dependency.dependencyNumber == triadDependency {
				// fmt.Printf("they are same!!!!!!!!\n")
				sameTriad := triad.Same(triads.GetElement(dependency.triad.Number()), i)
				triads.SetElement(i, &sameTriad)
				updateDep = false
			}
		}
		if _, ok := t.(*triad.AssignmentTriad); ok {
			operandTable[leftOperand.Hash()] = i
		}
		if updateDep {
			triadTable[t.Hash()] = triadDependencyTableItem{t, triadDependency}
		}
		// fmt.Printf("tables: %v,\n%v\n", operandTable, triadTable)
	}
}

func checkSameOperand(operand triad.Operand, triads triad.TriadList) (triad.Operand, bool) {
	if operand != nil {
		if link, ok := operand.(triad.LinkOperand); ok {
			linkedTriad := triads.GetElement(link.LinkTo)
			if same, ok := linkedTriad.(*triad.SameTriad); ok {
				return triad.Link(same.SameAs), true
			}
		}
	}
	return operand, false
}

func getOperandDep(operand triad.Operand, table dependencyTable, triads triad.TriadList) int {
	if operand != nil {
		if link, ok := operand.(triad.LinkOperand); ok {
			linkedTriad := triads.GetElement(link.LinkTo)
			// fmt.Printf("operand %v (%s)\n", linkedTriad, linkedTriad.Hash())
			triadDep, ok := table[linkedTriad.Hash()]
			if ok {
				return triadDep
			}
		}
		if dep, ok := table[operand.Hash()]; ok {
			return dep
		}
	}
	return 0
}
