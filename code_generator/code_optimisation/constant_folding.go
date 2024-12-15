package code_optimisation

import (
	"errors"

	"goodhumored/lr1_object_code_generator/code_generator/triad"
)

type constantTable = map[string]string

func foldConstants(triadList *triad.TriadList) {
	table := make(constantTable)
	for _, t := range triadList.Triads() {
		// fmt.Printf("table: %v\n", table)
		tryGettingValueAndUpdateTriadList(t, triadList, &table)
	}
}

func tryGettingValueAndUpdateTriadList(t triad.Triad, list *triad.TriadList, table *constantTable) (string, error) {
	switch t.(type) {
	case *triad.AssignmentTriad:
		checkAndUpdateRightOperand(t, list, table)
		if val, ok := tryGetTriadValue(t); ok {
			(*table)[t.Hash()] = val
		}
	case *triad.OrTriad, *triad.AndTriad, *triad.XorTriad, *triad.NotTriad:
		checkAndUpdateLeftOperand(t, list, table)
		// fmt.Printf("left operand value: %v\n", t.Left())
		checkAndUpdateRightOperand(t, list, table)
		// fmt.Printf("right operand value: %v\n", t.Right())
		if val, ok := tryGetTriadValue(t); ok {
			constTriad := triad.C(t.Number(), val)
			list.SetElement(t.Number(), &constTriad)
		}
	}
	// fmt.Printf("triad: %v\n", t)
	return "", errors.New("no value")
}

func checkAndUpdateLeftOperand(t triad.Triad, list *triad.TriadList, table *constantTable) {
	if strVal, ok := checkOperand(t.Left(), list); ok {
		t.SetLeft(triad.Id(strVal))
	}
	if value, ok := (*table)[t.Left().Hash()]; ok {
		t.SetLeft(triad.Id(value))
	}
}

func checkAndUpdateRightOperand(t triad.Triad, list *triad.TriadList, table *constantTable) {
	if t.Right() == nil {
		return
	}
	if strVal, ok := checkOperand(t.Right(), list); ok {
		t.SetRight(triad.Id(strVal))
	}
	if value, ok := (*table)[t.Right().Hash()]; ok {
		t.SetLeft(triad.Id(value))
	}
}

func checkOperand(operand triad.Operand, list *triad.TriadList) (string, bool) {
	if linkOperand, ok := operand.(triad.LinkOperand); ok {
		linkedTriad := list.GetElement(linkOperand.LinkTo)
		if constant, ok := linkedTriad.(*triad.ConstantTriad); ok {
			value, _ := constant.Value()
			strVal := value.(string)
			return strVal, true
		}
	}
	return "", false
}

func tryGetTriadValue(t triad.Triad) (string, bool) {
	if value, err := t.Value(); err == nil {
		strVal, ok := value.(string)
		return strVal, ok
	}
	return "", false
}
