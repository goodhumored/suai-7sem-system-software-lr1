package code_optimisation

import (
	"errors"
	"fmt"

	"goodhumored/lr1_object_code_generator/code_generator/triad"
)

type constantTable = map[string]string

// Функция реализующая алгоритм оптимизации методом свёртки констант
func foldConstants(triadList *triad.TriadList) {
	table := make(constantTable)
	for _, t := range triadList.Triads() {
		tryGettingValueAndUpdateTriadList(t, triadList, &table)
	}
}

// функция, которая пробует сосчитать значение триады и заменяет её на C(значение,)
// ссылки на другие C() триады функция заменяет на значения внутри них
func tryGettingValueAndUpdateTriadList(t triad.Triad, list *triad.TriadList, table *constantTable) (string, error) {
	switch t.(type) {
	case *triad.AssignmentTriad:
		checkAndUpdateRightOperand(t, list, table)
		if val, ok := tryGetTriadValue(t); ok {
			(*table)[t.Hash()] = val
		}
	case *triad.OrTriad, *triad.AndTriad, *triad.XorTriad, *triad.NotTriad:
		checkAndUpdateLeftOperand(t, list, table)
		checkAndUpdateRightOperand(t, list, table)
		if val, ok := tryGetTriadValue(t); ok {
			constTriad := triad.C(t.Number(), val)
			list.SetElement(t.Number(), &constTriad)
		}
	}
	return "", errors.New("no value")
}

// функция для проверки левого операнда на возможность получения значения
func checkAndUpdateLeftOperand(t triad.Triad, list *triad.TriadList, table *constantTable) {
	if strVal, ok := tryGettingOperandValue(t.Left(), list); ok {
		t.SetLeft(triad.Id(strVal))
	}
	if value, ok := (*table)[t.Left().Hash()]; ok {
		t.SetLeft(triad.Id(value))
	}
}

// функция для проверки правого операнда на возможность получения значения
func checkAndUpdateRightOperand(t triad.Triad, list *triad.TriadList, table *constantTable) {
	if t.Right() == nil {
		return
	}
	if strVal, ok := tryGettingOperandValue(t.Right(), list); ok {
		t.SetRight(triad.Id(strVal))
	}
	fmt.Printf("RIght operand hash: %s\n", t.Right().Hash())
	if value, ok := (*table)[t.Right().Hash()]; ok {
		fmt.Printf("have it in table\n")
		t.SetRight(triad.Id(value))
	}
}

// функция для проверки операнда на ссылание на константную триаду
func tryGettingOperandValue(operand triad.Operand, list *triad.TriadList) (string, bool) {
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

// функция пробует получить значение триады (например для триад бинарных операций),
// возвращает значение и то было ли это успешно
func tryGetTriadValue(t triad.Triad) (string, bool) {
	if value, err := t.Value(); err == nil {
		strVal, ok := value.(string)
		return strVal, ok
	}
	return "", false
}
