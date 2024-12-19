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

// функция оптимизации, реализующая алгоритм исключения лишних операций
func eliminateCommonSubexpression(triads *triad.TriadList) {
	operandTable := dependencyTable{}
	triadTable := map[string]triadDependencyTableItem{}
	for i, t := range triads.Triads() {
		updateDep := true
		leftOperand, ok := updateOperandIfLinkedToSameTriad(t.Left(), *triads)
		if ok {
			t.SetLeft(leftOperand)
		}
		rightOperand, ok := updateOperandIfLinkedToSameTriad(t.Right(), *triads)
		if ok {
			t.SetRight(rightOperand)
		}
		leftDependency := getOperandDependencyNumber(leftOperand, operandTable, *triads)
		rightDependency := getOperandDependencyNumber(rightOperand, operandTable, *triads)

		triadDependency := max(leftDependency, rightDependency) + 1
		if dependency, ok := triadTable[t.Hash()]; ok {
			if dependency.dependencyNumber == triadDependency {
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
	}
}

// функция которая проверяет, ссылается ли операнд на триаду Same, и возвращает новый операнд,
// ссылающийся напрямую на триаду, на которую ссылается триада Same
// Если операнд ссылается не ссылка или не ссылается на Same, то операнд возвращается без изменений
func updateOperandIfLinkedToSameTriad(operand triad.Operand, triads triad.TriadList) (triad.Operand, bool) {
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

// функция для получения числа зависимости операнда
func getOperandDependencyNumber(operand triad.Operand, table dependencyTable, triads triad.TriadList) int {
	if operand != nil {
		if link, ok := operand.(triad.LinkOperand); ok {
			linkedTriad := triads.GetElement(link.LinkTo)
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
