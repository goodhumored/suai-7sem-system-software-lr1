## ./code_generator/asm_8086_triad_translator/asm_8086_triad_translator.go
```go
package asm8086triadtranslator

import (
	"errors"
	"fmt"

	"goodhumored/lr1_object_code_generator/code_generator/triad"
)

// Транслятор внутреннего представления в объектный код asm 8086
type Asm8086TriadTranslator struct{}

// метод транслятора для перевода триад в результирующий код
func (t Asm8086TriadTranslator) TranslateTriads(triads triad.TriadList) (string, error) {
	code := ""
	mapKeys := []string{}
	for _, triad := range triads.Triads() {
		mapKeys = append(mapKeys, triad.Hash())
		triadCode, err := translateTriad(triad)
		if err != nil {
			return "", err
		}
		code += triadCode
	}

	return code, nil
}

// Функция для перевода триады в асм
func translateTriad(triadToTranslate triad.Triad) (string, error) {
	resultCode := ""
	switch triadToTranslate.(type) {
	case *triad.AssignmentTriad:
		return fmt.Sprintf("mov %s,%s\n", stringifyOperand(triadToTranslate.Left()), stringifyOperand(triadToTranslate.Right())), nil
	case *triad.AndTriad, *triad.OrTriad, *triad.XorTriad:
		act, _ := getActFromBinaryTriad(triadToTranslate)
		resultCode = fmt.Sprintf("mov ax,%s\n%s ax,%s\nmov tmp%d,ax\n", stringifyOperand(triadToTranslate.Left()), act, stringifyOperand(triadToTranslate.Right()), triadToTranslate.Number())
	case *triad.NotTriad:
		resultCode = fmt.Sprintf("mov ax,%s\nnot ax\n", stringifyOperand(triadToTranslate.Left()))
	case *triad.ConstantTriad, *triad.SameTriad:
	default:
		return "", fmt.Errorf("Неподдерживаемая триада %v\n", triadToTranslate)
	}
	return resultCode, nil
}

// Вспомогательная функция для получения строкового представления операнда для транслятора асм8086
func stringifyOperand(operand triad.Operand) string {
	if operand == nil {
		return ""
	}
	if linkOperand, ok := operand.(triad.LinkOperand); ok {
		return fmt.Sprintf("tmp%d", linkOperand.LinkTo)
	}
	if val, err := operand.Value(); err == nil {
		if strVal, ok := val.(string); ok {
			return strVal
		}
	}
	return ""
}

// функция получения строкового представления оператора из триады для транслятора в асм 8086
func getActFromBinaryTriad(t triad.Triad) (string, error) {
	switch t.(type) {
	case *triad.AndTriad:
		return "and", nil
	case *triad.OrTriad:
		return "or", nil
	case *triad.XorTriad:
		return "xor", nil
	}
	return "", errors.New("triad %t is not binary")
}
```

## ./code_generator/triad_mapper.go
```go
package codegenerator

import (
	"goodhumored/lr1_object_code_generator/code_generator/triad"
	"goodhumored/lr1_object_code_generator/syntax_analyzer/nonterminal"
	"goodhumored/lr1_object_code_generator/syntax_analyzer/parse_tree"
	"goodhumored/lr1_object_code_generator/token"
)

// функция преобразования дерева вывода в список триад
func MapParseTreeToTriadList(tree parse_tree.ParseTree) triad.TriadList {
	triads := triad.NewTriadList()
	_ = mapNodeToTriadList(*tree.Root, &triads)
	return triads
}

// функция преобразования узла дерева вывода в список триад
func mapNodeToTriadList(node parse_tree.Node, triads *triad.TriadList) triad.Operand {
	makeLink := true
	var outputOperand triad.Operand
	switch node.Symbol.GetName() {
	case token.IdentifierType.Name:
		outputOperand = triad.Id(node.Value)
		makeLink = false
	case nonterminal.Assignment.Name:
		mapAssignment(node, triads)
	case nonterminal.Binary.Name:
		mapBinary(node, triads)
	case nonterminal.Unary.Name:
		mapUnary(node, triads)
	default:
		for _, child := range node.Children {
			childOperand := mapNodeToTriadList(*child, triads)
			outputOperand = childOperand
			makeLink = false
		}
	}
	if makeLink && triads.Last() != nil {
		outputOperand = triad.Link(triads.Last())
	}
	return outputOperand
}

// функция преобразования бинарной операции в триаду
func mapBinary(node parse_tree.Node, triads *triad.TriadList) {
	operator := node.Children[1]
	operandNode1 := node.Children[0]
	operandNode2 := node.Children[2]

	operand1 := mapNodeToTriadList(*operandNode1, triads)
	operand2 := mapNodeToTriadList(*operandNode2, triads)

	var binaryTriad triad.Triad
	switch operator.Value {
	case "or":
		t := triad.Or(operand1, operand2, 0)
		binaryTriad = &t
	case "and":
		t := triad.And(operand1, operand2, 0)
		binaryTriad = &t
	case "xor":
		t := triad.Xor(operand1, operand2, 0)
		binaryTriad = &t
	}
	triads.Add(binaryTriad)
}

// функция преобразования операции присвоения в триаду
func mapAssignment(node parse_tree.Node, triads *triad.TriadList) {
	identifierOperandNode := node.Children[0]
	rightOperandNode := node.Children[2]
	identifierOperand := triad.Id(identifierOperandNode.Value)
	rightOperand := mapNodeToTriadList(*rightOperandNode, triads)
	assignmentTriad := triad.Assignment(identifierOperand, rightOperand, 0)
	triads.Add(&assignmentTriad)
}

// функция преобразования унарной операции в триаду
func mapUnary(node parse_tree.Node, triads *triad.TriadList) {
	operandNode := node.Children[2]
	operand := mapNodeToTriadList(*operandNode, triads)
	notTriad := triad.Not(operand, 0)
	triads.Add(&notTriad)
}
```

## ./code_generator/triad/and_triad.go
```go
package triad

import (
	"fmt"
)

type AndTriad struct {
	LogicTriad
}

func (t AndTriad) Hash() string {
	return fmt.Sprintf("and_%s_%s", t.left.Hash(), t.right.Hash())
}

func (t AndTriad) String() string {
	return fmt.Sprintf("and(%s,%s)", t.left.String(), t.right.String())
}

func And(left Operand, right Operand, number int) AndTriad {
	return AndTriad{
		LogicTriad: Logic(number, left, right, func(left int, right int) int {
			return left & right
		}),
	}
}
```

## ./code_generator/triad/same_triad.go
```go
package triad

import "fmt"

type SameTriad struct {
	baseTriad
	SameAs Triad
}

func (t SameTriad) Value() (any, error) {
	return t.left.Value()
}

func (t SameTriad) Hash() string {
	return t.left.Hash()
}

func (t SameTriad) String() string {
	return fmt.Sprintf("Same(%d,)", t.SameAs.Number())
}

func Same(triad Triad, number int) SameTriad {
	return SameTriad{
		baseTriad: baseTriad{number: number, left: triad, right: nil},
		SameAs:    triad,
	}
}
```

## ./code_generator/triad/assignment_triad.go
```go
package triad

import (
	"fmt"
)

type AssignmentTriad struct {
	baseTriad
}

func (t AssignmentTriad) Hash() string {
	return t.left.String()
}

func (t AssignmentTriad) String() string {
	return fmt.Sprintf(":=(%s,%s)", t.left.String(), t.right.String())
}

func (t AssignmentTriad) Value() (any, error) {
	return t.right.Value()
}

func Assignment(left Operand, right Operand, number int) AssignmentTriad {
	return AssignmentTriad{
		baseTriad{left: left, right: right, number: number},
	}
}
```

## ./code_generator/triad/id_operand.go
```go
package triad

type IdOperand struct{ name string }

func (o IdOperand) Hash() string {
	return o.name
}

func (o IdOperand) String() string {
	return o.name
}

func (o IdOperand) Value() (any, error) {
	return o.name, nil
}

func Id(name string) IdOperand {
	return IdOperand{
		name,
	}
}
```

## ./code_generator/triad/logic_triad.go
```go
package triad

import (
	"errors"
	"strconv"
)

type LogicTriad struct {
	baseTriad
	operation func(left int, right int) int
}

func (t LogicTriad) Value() (any, error) {
	leftIntVal, ok := parseOperand(t.left)
	if !ok {
		return nil, errors.New("failed parsing left")
	}
	rightIntVal, ok := parseOperand(t.right)
	if !ok {
		return nil, errors.New("failed parsing right")
	}
	value := t.operation(leftIntVal, rightIntVal)
	return strconv.Itoa(value), nil
}

func parseOperand(operand Operand) (int, bool) {
	if _, isId := operand.(IdOperand); !isId {
		return 0, false
	}
	val, err := operand.Value()
	if err != nil {
		return 0, false
	}
	strVal, ok := val.(string)
	if !ok {
		return 0, false
	}
	intVal, err := convertStrToNumber(strVal)
	if err != nil {
		return 0, false
	}
	return int(intVal), true
}

func convertStrToNumber(str string) (int64, error) {
	base := 10
	if len(str) > 1 && str[1] == 'x' {
		str = str[2:]
		base = 16
	}
	return strconv.ParseInt(str, base, 32)
}

func Logic(number int, left Operand, right Operand, operation func(left int, right int) int) LogicTriad {
	return LogicTriad{
		baseTriad: baseTriad{number: number, left: left, right: right},
		operation: operation,
	}
}
```

## ./code_generator/triad/link_operand.go
```go
package triad

import (
	"fmt"
)

type LinkOperand struct{ LinkTo int }

func (o LinkOperand) Hash() string {
	return fmt.Sprintf("^%d", o.LinkTo)
}

func (o LinkOperand) String() string {
	return fmt.Sprintf("^%d", o.LinkTo)
}

func (o LinkOperand) Value() (any, error) {
	return o.LinkTo, nil
}

func Link(triad Triad) LinkOperand {
	return LinkOperand{
		LinkTo: triad.Number(),
	}
}
```

## ./code_generator/triad/constant_triad.go
```go
package triad

import "fmt"

type ConstantTriad struct {
	baseTriad
	value any
}

func (t ConstantTriad) Value() (any, error) {
	return t.value, nil
}

func (t ConstantTriad) String() string {
	return fmt.Sprintf("C(%s,)", t.value.(string))
}

func (t ConstantTriad) Hash() string {
	return fmt.Sprintf("c_%s", t.value)
}

func C[T any](number int, value T) ConstantTriad {
	return ConstantTriad{
		baseTriad: baseTriad{number: number, left: nil, right: nil},
		value:     value,
	}
}
```

## ./code_generator/triad/operand.go
```go
package triad

type Operand interface {
	Value() (any, error)
	Hash() string
	String() string
}
```

## ./code_generator/triad/triad_list.go
```go
package triad

import "fmt"

type TriadList struct {
	list   []Triad
	length int
}

func (l *TriadList) Add(triad Triad) {
	triad.SetNumber(l.length)
	l.list = append(l.list, triad)
	l.length++
}

func (l TriadList) Triads() []Triad {
	return l.list
}

func (l TriadList) Print() {
	for _, triad := range l.Triads() {
		fmt.Printf("%d)%s\n", triad.Number(), triad.String())
	}
}

func (l TriadList) GetElement(n int) Triad {
	return l.list[n]
}

func (l *TriadList) SetElement(n int, triad Triad) {
	fmt.Printf("setting triad %d)%v to %v\n", n, l.list[n], triad)
	l.list[n] = triad
}

func (l TriadList) Remove(number int) {
	for i := number; i < l.length; i++ {
		triadEl := l.list[i]
		triadEl.SetNumber(triadEl.Number() - 1)
		if operand, ok := triadEl.Left().(LinkOperand); ok {
			operand.LinkTo--
		}
		if operand, ok := triadEl.Right().(LinkOperand); ok {
			operand.LinkTo--
		}
	}
	l.list = append(l.list[:number], l.list[number+1:]...)
	l.length--
}

func (l TriadList) Last() Triad {
	if l.length > 0 {
		return l.list[l.length-1]
	}
	return nil
}

func NewTriadList() TriadList {
	return TriadList{list: []Triad{}, length: 0}
}
```

## ./code_generator/triad/not_triad.go
```go
package triad

import (
	"errors"
	"fmt"
	"strconv"
)

type NotTriad struct {
	baseTriad
}

func (t NotTriad) String() string {
	return fmt.Sprintf("not(%s,)", t.left.String())
}

func (t NotTriad) Hash() string {
	return fmt.Sprintf("not_%s", t.left.Hash())
}

func (t NotTriad) Value() (any, error) {
	if leftVal, err := t.left.Value(); err == nil {
		if strVal, ok := leftVal.(string); ok {
			strVal = strVal[2:]
			intVal, err := strconv.ParseInt(strVal, 16, 32)
			if err != nil {
				return nil, err
			}
			return strconv.Itoa(int(^intVal)), nil
		}
	}
	return 0, errors.New("no value")
}

func Not(operand Operand, number int) NotTriad {
	return NotTriad{
		baseTriad{
			left:   operand,
			number: number,
		},
	}
}
```

## ./code_generator/triad/or_triad.go
```go
package triad

import (
	"fmt"
)

type OrTriad struct {
	LogicTriad
}

func (t OrTriad) Hash() string {
	return fmt.Sprintf("or_%s_%s", t.left.Hash(), t.right.Hash())
}

func (t OrTriad) String() string {
	return fmt.Sprintf("or(%s,%s)", t.left.String(), t.right.String())
}

func Or(left Operand, right Operand, number int) OrTriad {
	return OrTriad{
		LogicTriad: Logic(number, left, right, func(left int, right int) int {
			return left | right
		}),
	}
}
```

## ./code_generator/triad/triad.go
```go
package triad

type Triad interface {
	Operand
	Number() int
	SetNumber(nubmer int)
	Left() Operand
	SetLeft(Operand)
	SetRight(Operand)
	Right() Operand
	Hash() string
}

type baseTriad struct {
	number int
	left   Operand
	right  Operand
}

func (t baseTriad) Number() int {
	return t.number
}

func (t *baseTriad) SetNumber(number int) {
	t.number = number
}

func (t *baseTriad) SetLeft(newLeft Operand) {
	t.left = newLeft
}

func (t *baseTriad) SetRight(newRight Operand) {
	t.right = newRight
}

func (t baseTriad) Left() Operand {
	return t.left
}

func (t baseTriad) Right() Operand {
	return t.right
}

func (t baseTriad) Hash() string {
	hash := ""
	if t.left != nil {
		hash += t.left.Hash()
	}
	if t.right != nil {
		hash += "_" + t.right.Hash()
	}
	return hash
}
```

## ./code_generator/triad/xor_triad.go
```go
package triad

import (
	"fmt"
)

type XorTriad struct {
	LogicTriad
}

func (t XorTriad) Hash() string {
	return fmt.Sprintf("xor_%s_%s", t.left.Hash(), t.right.Hash())
}

func (t XorTriad) String() string {
	return fmt.Sprintf("xor(%s,%s)", t.left.String(), t.right.String())
}

func Xor(left Operand, right Operand, number int) XorTriad {
	return XorTriad{
		LogicTriad: Logic(number, left, right, func(left int, right int) int {
			return left ^ right
		}),
	}
}
```

## ./code_generator/code_generator.go
```go
package codegenerator

import (
	"fmt"
	"strings"

	asm8086triadtranslator "goodhumored/lr1_object_code_generator/code_generator/asm_8086_triad_translator"
	code_optimisation "goodhumored/lr1_object_code_generator/code_generator/code_optimisation"
	"goodhumored/lr1_object_code_generator/syntax_analyzer/parse_tree"
)

// Функция генерации обектного кода на основе дерева вывода
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
```

## ./code_generator/code_optimisation/constant_folding.go
```go
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
```

## ./code_generator/code_optimisation/common_subexpression_elimination.go
```go
package code_optimisation

import (
	"fmt"

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
				fmt.Printf("triad %s has same previous occusion\n", t.Hash())
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
```

## ./code_generator/code_optimisation/code_optimiser.go
```go
package code_optimisation

import (
	"fmt"

	"goodhumored/lr1_object_code_generator/code_generator/triad"
)

// Функция оптимизирующая входной список триад
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
```

## ./code_generator/triad_translator.go
```go
package codegenerator

import "goodhumored/lr1_object_code_generator/code_generator/triad"

// интерфейс переводчика триад
type TriadTranslator interface {
	TranslateTriads(triad.TriadList) (string, error)
}
```

## ./token_analyzer/token_patterns.go
```go
package token_analyzer

import (
	"regexp"

	"goodhumored/lr1_object_code_generator/token"
)

// Вспомогатеьлная структура для установки соответствия шаблонов лексем
// с их фабричными функциями
type TokenPattern struct {
	Pattern *regexp.Regexp
	Type    func(string, token.Position) token.Token
}

// Массив соответствий шаблонов лексем
var tokenPatterns = []TokenPattern{
	{regex("or"), token.Or},
	{regex("xor"), token.Xor},
	{regex("and"), token.And},
	{regex("not"), token.Not},
	{regex("(0x|[0-9$])[0-9a-fA-F]+"), token.Identifier},
	{regex("[a-zA-Z][a-zA-Z0-9]+"), token.Identifier},
	{regex(":="), token.Assignment},
	{regex("#.*"), token.Comment},
	{regex("[(]"), token.LeftParenth},
	{regex("[)]"), token.RightParenth},
	{regex(";"), token.Delimiter},
}

// вспомогательная функция создающая объект регулярного выражения
// добавляющая в начале шаблона признак начала строки
func regex(pattern string) *regexp.Regexp {
	return regexp.MustCompile("^" + pattern)
}
```

## ./token_analyzer/token_analyzer.go
```go
package token_analyzer

import (
	"strings"

	"goodhumored/lr1_object_code_generator/token"
	"goodhumored/lr1_object_code_generator/token_table"
)

// Распознаёт токены в данной строке построчно и записывает в таблицу
func RecogniseTokens(source string) *token_table.TokenTable {
	tokenTable := &token_table.TokenTable{}
	tokenTable.Add(token.Start)
	for _, line := range strings.Split(source, "\n") {
		recogniseTokensLine(line, tokenTable)
	}
	tokenTable.Add(token.EOF)
	return tokenTable
}

// Распознаёт лексемы в данной строке и записывает в таблицу
func recogniseTokensLine(line string, tokenTable *token_table.TokenTable) {
	for {
		line = strings.Trim(line, " ") // обрезаем пробельные символы в строке
		if len(line) == 0 {            // если строка пустая - завершаем обработку строки
			return
		}
		nextToken := getNextToken(line)      // ищем очередную лексему
		tokenTable.Add(nextToken)            // добавляем лексему в таблицу
		line = line[nextToken.Position.End:] // вырезаем обработанную часть
	}
}

// Ищет очередную лексему в строке
func getNextToken(str string) token.Token {
	// проходим по всем шаблонам лексем
	for _, tokenPattern := range tokenPatterns {
		res := tokenPattern.Pattern.FindStringIndex(str)
		if res != nil {
			return tokenPattern.Type(str[res[0]:res[1]], token.Position{res[0], res[1]})
		}
	}
	return token.Error(str[0:1], token.Position{0, 1})
}
```

## ./main.go
```go
package main

import (
	"fmt"
	"os"

	codegenerator "goodhumored/lr1_object_code_generator/code_generator"
	"goodhumored/lr1_object_code_generator/syntax_analyzer"
	"goodhumored/lr1_object_code_generator/token_analyzer"
)

func main() {
	source := getInput("./input-simple.txt") // читаем файл

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
```

## ./token/token.go
```go
package token

import "fmt"

// Структура Position представляет положение лексемы в строке
type Position struct {
	Start int
	End   int
}

// Структура Token представляет лексему с ее типом и значением
type Token struct {
	Type     TokenType // Тип
	value    string    // Значение
	Position Position  // Положение лексемы
}

// Функция получения имени токена, для соответствия интерфейсу символа
func (token Token) GetName() string {
	return token.Type.GetName()
}

func (token Token) Value() string {
	return token.value
}

// Фабричная функция для токенов, возвращающая замкнутую лямбда функцию для создания токена определённого типа
func tokenFactory(tokenType TokenType) func(string, Position) Token {
	return func(value string, position Position) Token {
		return Token{
			value:    value,
			Type:     tokenType,
			Position: position,
		}
	}
}

// Функция определяющая как токен преобразуется в строку
func (token Token) String() string {
	return fmt.Sprintf("%s (%s)", token.Type, token.Value())
}

// Функции создания лексем определённых типов
var (
	Delimiter    = tokenFactory(DelimiterType)           // Разделиитель
	Identifier   = tokenFactory(IdentifierType)          // Идентификатор
	Assignment   = tokenFactory(AssignmentType)          // Присваивание
	And          = tokenFactory(AndType)                 // И
	Or           = tokenFactory(OrType)                  // Или
	Xor          = tokenFactory(XorType)                 // Исключающее или
	Not          = tokenFactory(NotType)                 // Не
	LeftParenth  = tokenFactory(LeftParenthType)         // Левая скобка
	RightParenth = tokenFactory(RightParenthType)        // Правая скобка
	Error        = tokenFactory(ErrorType)               // Ошибка
	Comment      = tokenFactory(CommentType)             // Комментарий
	Start        = Token{StartType, "", Position{0, 0}}  // Начало строки
	EOF          = Token{EOFType, "EOF", Position{0, 0}} // Конец
)
```

## ./token/token_type.go
```go
package token

type TokenType struct {
	Name string
}

func (tokenType TokenType) GetName() string {
	return tokenType.Name
}

var (
	DelimiterType    = TokenType{"delimiter"}         // Разделитель
	IdentifierType   = TokenType{"identifier"}        // Идентификатор
	ConstantType     = TokenType{"const_number"}      // Шестнадцатиричное число
	AssignmentType   = TokenType{"assignment"}        // Присваивание
	AndType          = TokenType{"and"}               // and
	OrType           = TokenType{"or"}                // or
	XorType          = TokenType{"xor"}               // xor
	NotType          = TokenType{"not"}               // not
	LeftParenthType  = TokenType{"left_parentheses"}  // Скобки
	RightParenthType = TokenType{"right_parentheses"} // Скобки
	ErrorType        = TokenType{"error"}             // Ошибка
	CommentType      = TokenType{"comment"}           // Комментарий
	StartType        = TokenType{"start"}             // Начало
	EOFType          = TokenType{"EOF"}               // Конец
)
```

## ./rules_table.go
```go
package main

import (
	"goodhumored/lr1_object_code_generator/syntax_analyzer/nonterminal"
	"goodhumored/lr1_object_code_generator/syntax_analyzer/rule"
	"goodhumored/lr1_object_code_generator/token"
)

func or(symbols ...rule.Symbol) []rule.Symbol {
	return symbols
}

var (
	valueSymbols           = or(nonterminal.Binary, nonterminal.Unary, token.IdentifierType, token.ConstantType, nonterminal.Parenthesis, nonterminal.Value)
	binaryOperatorsSymbols = or(token.AndType, token.OrType, token.XorType)
)

// Правила грамматики
var assignmentRule = rule.Rule{
	Left:  nonterminal.Assignment,
	Right: [][]rule.Symbol{or(token.IdentifierType), or(token.AssignmentType), valueSymbols, or(token.DelimiterType)},
}

var binaryRule = rule.Rule{
	Left:  nonterminal.Binary,
	Right: [][]rule.Symbol{valueSymbols, binaryOperatorsSymbols, valueSymbols},
}

var unaryRule = rule.Rule{
	Left:  nonterminal.Unary,
	Right: [][]rule.Symbol{or(token.NotType), or(token.LeftParenthType), valueSymbols, or(token.RightParenthType)},
}

var parenthesisRule = rule.Rule{
	Left:  nonterminal.Parenthesis,
	Right: [][]rule.Symbol{or(token.LeftParenthType), valueSymbols, or(token.RightParenthType)},
}

var identifierRule = rule.Rule{
	Left:  nonterminal.Value,
	Right: [][]rule.Symbol{valueSymbols},
}

var rootRule = rule.Rule{
	Left:  nonterminal.Root,
	Right: [][]rule.Symbol{or(token.StartType), or(nonterminal.Assignment), or(token.EOFType)},
}

var rulesTable = rule.RuleTable{Rules: []rule.Rule{
	unaryRule,
	parenthesisRule,
	binaryRule,
	assignmentRule,
	rootRule,
	identifierRule,
}}
```

## ./token_table/token_table.go
```go
package token_table

import (
	"fmt"
	"strings"

	"goodhumored/lr1_object_code_generator/token"
)

// Таблица лексем
type TokenTable struct {
	tokens []token.Token
}

// Метод добавления лексемы в таблицу
func (tt *TokenTable) Add(token token.Token) {
	tt.tokens = append(tt.tokens, token)
}

// Метод получения списка найденных лексем
func (tt TokenTable) GetTokens() []token.Token {
	return tt.tokens
}

// Вспомогательная функция для вывода таблицы
func (tt *TokenTable) Print() {
	errors := tt.GetErrors()
	if len(errors) > 0 {
		errorMsg := ""
		for _, error := range errors {
			errorMsg += fmt.Sprintf("Неизвестный символ: %s \n", error.Value())
		}
		fmt.Println(fmt.Errorf(errorMsg))
	}
	fmt.Println(tt.String())
}

// Вспомогательная функция для генерации строки с таблицей лексем
func (tt *TokenTable) String() string {
	if len(tt.tokens) == 0 {
		return "Ни одного токена не найдено"
	}

	// Определяем максимальную ширину столбца
	maxTypeLen := len("Тип")
	maxValueLen := len("Значение")
	for _, token := range tt.tokens {
		if len(token.Type.Name) > maxTypeLen {
			maxTypeLen = len(token.Type.Name)
		}
		if len(token.Value()) > maxValueLen {
			maxValueLen = len(token.Value())
		}
	}

	// создаем шапку и рамки
	header := fmt.Sprintf("| %-*s | %-*s |", maxTypeLen, "Тип", maxValueLen, "Значение")
	border := fmt.Sprintf("+-%s-+-%s-+", strings.Repeat("-", maxTypeLen), strings.Repeat("-", maxValueLen))

	// Собираем таблицу
	res := border + "\n" + header + "\n" + border + "\n"
	for _, token := range tt.tokens {
		res += fmt.Sprintf("| %-*s | %-*s |\n", maxTypeLen, token.GetName(), maxValueLen, token.Value())
	}
	res += border

	return res
}

// Функция возвращающая все ошибки в таблице
func (tt TokenTable) GetErrors() []token.Token {
	tokens := []token.Token{}
	for _, recognisedToken := range tt.tokens {
		if recognisedToken.Type == token.ErrorType {
			tokens = append(tokens, recognisedToken)
		}
	}
	return tokens
}
```

## ./precedence_matrix.go
```go
package main

import (
	"goodhumored/lr1_object_code_generator/syntax_analyzer/precedence"
	"goodhumored/lr1_object_code_generator/token"
)

// Матрица предшествования
var precedenceMatrix = precedence.Matrix{
	token.IdentifierType:   precedence.Row{token.AssignmentType: precedence.Eq, token.RightParenthType: precedence.Gt, token.OrType: precedence.Gt, token.XorType: precedence.Gt, token.AndType: precedence.Gt, token.DelimiterType: precedence.Gt},
	token.AssignmentType:   precedence.Row{token.IdentifierType: precedence.Lt, token.LeftParenthType: precedence.Lt, token.RightParenthType: precedence.Lt, token.NotType: precedence.Lt, token.OrType: precedence.Lt, token.XorType: precedence.Lt, token.AndType: precedence.Lt, token.DelimiterType: precedence.Eq},
	token.LeftParenthType:  precedence.Row{token.IdentifierType: precedence.Lt, token.LeftParenthType: precedence.Lt, token.RightParenthType: precedence.Eq, token.NotType: precedence.Lt, token.OrType: precedence.Lt, token.XorType: precedence.Lt, token.AndType: precedence.Lt},
	token.RightParenthType: precedence.Row{token.RightParenthType: precedence.Gt, token.OrType: precedence.Gt, token.XorType: precedence.Gt, token.AndType: precedence.Gt, token.DelimiterType: precedence.Gt},
	token.NotType:          precedence.Row{token.LeftParenthType: precedence.Lt},
	token.OrType:           precedence.Row{token.IdentifierType: precedence.Lt, token.LeftParenthType: precedence.Lt, token.RightParenthType: precedence.Gt, token.NotType: precedence.Lt, token.OrType: precedence.Gt, token.XorType: precedence.Gt, token.AndType: precedence.Lt, token.DelimiterType: precedence.Gt},
	token.XorType:          precedence.Row{token.IdentifierType: precedence.Lt, token.LeftParenthType: precedence.Lt, token.RightParenthType: precedence.Gt, token.NotType: precedence.Lt, token.OrType: precedence.Gt, token.XorType: precedence.Gt, token.AndType: precedence.Lt, token.DelimiterType: precedence.Gt},
	token.AndType:          precedence.Row{token.IdentifierType: precedence.Lt, token.LeftParenthType: precedence.Lt, token.RightParenthType: precedence.Gt, token.NotType: precedence.Lt, token.OrType: precedence.Gt, token.XorType: precedence.Gt, token.AndType: precedence.Gt, token.DelimiterType: precedence.Gt},
	token.DelimiterType:    precedence.Row{token.IdentifierType: precedence.Gt},
}
```

## ./syntax_analyzer/parse_tree/node.go
```go
package parse_tree

import (
	"fmt"

	"goodhumored/lr1_object_code_generator/syntax_analyzer/rule"
)

// Узел дерева вывода
type Node struct {
	Symbol   rule.Symbol
	Value    string
	Children []*Node
}

// Вспомогательная функция для создания пустого узла
func CreateNode(s rule.Symbol) Node {
	return Node{Symbol: s, Children: []*Node{}, Value: ""}
}

// Метод, добавляющий дочерний узел
func (n *Node) AddChild(child *Node) {
	n.Children = append(n.Children, child)
}

// Метод свёртки узла дерева
func (node *Node) Reduce(rule rule.Rule) bool {
	// Если не можем применить правило к текущему узлу - уходим
	if !node.CanApplyRule(rule) {
		return false
	}
	// считаем разницу длин правой части правила и детей узла
	lenDiff := len(node.Children) - len(rule.Right)

	// копируем слайс с нужными нам узлами, которые собираемся заменять
	nodes := make([]*Node, len(rule.Right))
	copy(nodes, node.Children[lenDiff:])

	// перезаписываем дочерние узлы узла
	node.Children = append(node.Children[:lenDiff], &Node{Symbol: rule.Left, Children: nodes, Value: ""})
	return true
}

// Функция проверки возможности применения правила к дочерним узлам узла
func (node Node) CanApplyRule(ruleToCheck rule.Rule) bool {
	childrenSymbols := []rule.Symbol{}
	for _, child := range node.Children {
		childrenSymbols = append(childrenSymbols, child.Symbol)
	}
	return rule.IsApplyable(ruleToCheck.Right, childrenSymbols)
}

// Метод для рекурсивного вывода узлов дерева в консоль
func (node *Node) Print(prefix string, isTail bool) {
	// Выводим символ узла с отступом
	var branch, prefixSuffix string
	if isTail {
		prefixSuffix = "    "
		branch = "└── "
	} else {
		branch = "├── "
		prefixSuffix = "│   "
	}
	fmt.Println(prefix + branch + node.Symbol.GetName() + " (" + node.Value + ")")

	// Рекурсивно выводим дочерние узлы
	for i := 0; i < len(node.Children)-1; i++ {
		node.Children[i].Print(prefix+prefixSuffix, false)
	}
	if len(node.Children) > 0 {
		node.Children[len(node.Children)-1].Print(prefix+prefixSuffix, true)
	}
}
```

## ./syntax_analyzer/parse_tree/parse_tree.go
```go
package parse_tree

import (
	"fmt"

	"goodhumored/lr1_object_code_generator/syntax_analyzer/rule"
)

// Дерево вывода
type ParseTree struct {
	Root *Node
}

// Метод добавления узлов в дерево
func (tree *ParseTree) AddNode(node *Node) {
	tree.Root.AddChild(node)
}

// Метод для свёртки дерева по правилу
func (tree *ParseTree) Reduce(rule rule.Rule) {
	fmt.Printf("Применяем правило %s к дереву\n", rule)
	if tree.Root.Reduce(rule) {
		fmt.Printf("Успешно применено\n")
	} else {
		fmt.Printf("Правило %s применить не удалось\n", rule)
	}
}

// Метод для вывода дерева
func (tree ParseTree) Print() {
	tree.Root.Print("", true)
}
```

## ./syntax_analyzer/nonterminal/nonterminal.go
```go
package nonterminal

// Структура представляющая нетерминалы
type NonTerminal struct {
	Name string
}

// Метод для соответствия нетерменалов интерфейсу символ
func (nt NonTerminal) GetName() string {
	return nt.Name
}

func (nt NonTerminal) Value() string {
	return ""
}

var (
	E           = NonTerminal{"E"}                // Стандартный нетерминал
	Parenthesis = NonTerminal{"PARENTHESIS"}      // Стандартный нетерминал
	Assignment  = NonTerminal{"ASSIGNMENT"}       // Стандартный нетерминал
	Binary      = NonTerminal{"BINARY_OPERATION"} // Стандартный нетерминал
	Value       = NonTerminal{"VALUE"}            // Стандартный нетерминал
	Unary       = NonTerminal{"UNARY_OPERATION"}  // Стандартный нетерминал
	Root        = NonTerminal{"/"}                // Корневой нетерминал
)
```

## ./syntax_analyzer/symbol_stack.go
```go
package syntax_analyzer

import (
	"fmt"

	"goodhumored/lr1_object_code_generator/syntax_analyzer/rule"
	"goodhumored/lr1_object_code_generator/token"
)

// Стек символов
type symbolStack []rule.Symbol

// Добавление символа в стек
func (s symbolStack) Push(e rule.Symbol) symbolStack {
	return append(s, e)
}

// Просмотр верхнего элемента стека
func (s symbolStack) Peek() rule.Symbol {
	length := len(s)
	if length == 0 {
		return nil
	}
	return s[length-1]
}

// Просмотр n-ного элемента стека
func (s symbolStack) PeekN(n int) rule.Symbol {
	length := len(s)
	if length == 0 {
		return nil
	}
	return s[length-n-1]
}

// Поиск ближайшего к вершине терминала в стеке
func (s symbolStack) PeekTopTerminal() *token.Token {
	for i := range s {
		symbol := s.PeekN(i)
		if token, ok := symbol.(token.Token); ok {
			return &token
		}
	}
	return nil
}

// Вспомогательный метод преобразования стека символов в строку
func (s symbolStack) String() string {
	str := ""
	for _, i := range s {
		str += fmt.Sprintf("%s ", i.GetName())
	}
	return str
}

// Вспомогательный метод вывода стека символов
func (s symbolStack) Print() {
	fmt.Print(s.String())
}
```

## ./syntax_analyzer/rule/rule.go
```go
package rule

import (
	"fmt"

	"goodhumored/lr1_object_code_generator/syntax_analyzer/nonterminal"
)

// Интерфейс для представления символа
type Symbol interface {
	GetName() string
}

// Правило
type Rule struct {
	Left  nonterminal.NonTerminal
	Right [][]Symbol
}

// Метод получения строки из правила
func (r Rule) String() string {
	return fmt.Sprintf("%s -> %s", r.Left.GetName(), r.Right)
}
```

## ./syntax_analyzer/rule/rule_table.go
```go
package rule

// Таблица правил
type RuleTable struct {
	Rules []Rule
}

// Метод поиска правила по правой части
func (ruleTable RuleTable) GetRuleByRightSide(tokenTypes []Symbol) *Rule {
	for _, rule := range ruleTable.Rules {
		if IsApplyable(rule.Right, tokenTypes) {
			return &rule
		}
	}
	return nil
}

// Проверка на применимость правила к целевым символам
func IsApplyable(ruleSymbols [][]Symbol, targetSymbols []Symbol) bool {
	// Проверяем длины
	lenDiff := len(targetSymbols) - len(ruleSymbols)
	if lenDiff < 0 {
		return false
	}
	// Сравниваем последние символы цепочки символов и символы правила
	for i, ruleSymbol := range ruleSymbols {
		if !ContainsRule(ruleSymbol, targetSymbols[i+lenDiff]) {
			return false
		}
	}
	return true
}

func ContainsRule(arr []Symbol, el Symbol) bool {
	elName := el.GetName()
	for _, e := range arr {
		if e.GetName() == elName {
			return true
		}
	}
	return false
}
```

## ./syntax_analyzer/precedence/precedence_type.go
```go
package precedence

// Тип предшествования
type PrecedenceType struct {
	Name string
}

// Типы предшествования
var (
	Lt        = PrecedenceType{"<"} // Предшествует
	Eq        = PrecedenceType{"="} // Составляет основу
	Gt        = PrecedenceType{">"} // Следует
	Undefined = PrecedenceType{"-"} // Неопределено
)
```

## ./syntax_analyzer/precedence/precenence_matrix.go
```go
package precedence

import (
	"goodhumored/lr1_object_code_generator/token"
)

type Row = map[token.TokenType]PrecedenceType

// Матрица предшествования
type Matrix map[token.TokenType]Row

// Метод для поиска типа предшествования для двух терминалов
func (matrix Matrix) GetPrecedence(left, right token.TokenType) PrecedenceType {
	// Если левый символ - начало файла, возвращаем предшествие
	if left == token.StartType {
		return Lt
	}
	// Если правый символ - конец файла, возвращаем следствие
	if right == token.EOFType {
		return Gt
	}
	// Если находится - возвращаем
	if val, ok := matrix[left]; ok {
		if precedence, ok := val[right]; ok {
			return precedence
		}
	}
	// Если не находится - возвращаем неопределённость
	return Undefined
}
```

## ./syntax_analyzer/syntax_analyzer.go
```go
package syntax_analyzer

import (
	"errors"
	"fmt"

	"goodhumored/lr1_object_code_generator/syntax_analyzer/nonterminal"
	"goodhumored/lr1_object_code_generator/syntax_analyzer/parse_tree"
	"goodhumored/lr1_object_code_generator/syntax_analyzer/precedence"
	"goodhumored/lr1_object_code_generator/syntax_analyzer/rule"
	"goodhumored/lr1_object_code_generator/token"
	"goodhumored/lr1_object_code_generator/token_table"
)

// Функция для анализа синтаксиса, принимает таблицу токенов, список правил и матрицу предшествования
func AnalyzeSyntax(ruleTable rule.RuleTable, tokenTable token_table.TokenTable, matrix precedence.Matrix) (parse_tree.ParseTree, error) {
	// Создаём дерево
	rootNode := parse_tree.CreateNode(nonterminal.Root)
	tree := parse_tree.ParseTree{Root: &rootNode}
	// Получаем лексемы из таблицы
	tokens := tokenTable.GetTokens()
	tokenIndex := 1
	// Создаём стек
	stack := symbolStack{tokens[0]}

	for {
		// Берём ближайший к вершине терминал
		stackTerminal := stack.PeekTopTerminal()
		// Берём текущий символ входной строки
		if len(tokens) <= tokenIndex {
			return tree, errors.New("Токены закончились, до конца свернуть не удалось")
		}
		inputToken := tokens[tokenIndex]
		// Если строка принята, значит возвращаем дерево вывода
		if isInputAccepted(inputToken, stack) {
			return tree, nil
		}
		// Если комментарий - пропускаем
		if inputToken.Type == token.CommentType {
			tokenIndex += 1
			continue
		}

		fmt.Printf("Лексема: '%s' \n", tokens[tokenIndex].Value())
		fmt.Printf("Стек: %s \n", stack)

		// Получаем предшествование из матрицы
		prec := matrix.GetPrecedence(stackTerminal.Type, inputToken.Type)

		// Если предшествование или =, тогда сдвигаем
		if prec == precedence.Lt || prec == precedence.Eq {
			print("Сдвигаем\n")
			node := &parse_tree.Node{Value: inputToken.Value(), Symbol: inputToken, Children: []*parse_tree.Node{}}
			tree.AddNode(node) // Добавляем узел в дерево
			stack = stack.Push(inputToken)
			tokenIndex += 1
		} else if prec == precedence.Gt { // Иначе сворачиваем
			print("Сворачиваем\n")
			// сворачиваем стек
			newStack, rule, err := reduce(stack, ruleTable)
			if err != nil {
				return tree, err
			}
			stack = newStack
			// сворачиваем дерево
			tree.Reduce(*rule)
		} else {
			// Если предшествование не определено - выдаем ошибку
			return tree, fmt.Errorf("Ошибка в синтексе, неожиданное сочетание символов %s и %s (%d)", stackTerminal.GetName(), inputToken.GetName(), inputToken.Position.End)
		}
		println("==============")
	}
}

// Проверка на завершённость
func isInputAccepted(currentToken token.Token, stack symbolStack) bool {
	nextTerminal := stack.PeekTopTerminal()
	nextSymbol := stack.Peek()
	return currentToken.Type == token.EOFType && // Если дошли до конца строки
		nextTerminal != nil &&
		nextTerminal.Type == token.Start.Type && // Если ближайший терминал в стеке - начало строки
		nextSymbol != nil &&
		nextSymbol == nonterminal.Assignment // А на вершине строки - целевой символ
}

// Функция свёртки стека
func reduce(stack symbolStack, ruleTable rule.RuleTable) (symbolStack, *rule.Rule, error) {
	for {
		// Если есть применимое к стеку правило
		if rule := ruleTable.GetRuleByRightSide(stack); rule != nil {
			fmt.Printf("Нашлось правило: %v, пушим %s в стек\n", rule, rule.Left)
			// обновляем стек
			stack = append(stack[:len(stack)-len(rule.Right)], rule.Left)
			return stack, rule, nil
		} else {
			// Если нет выдаем ошибку
			return stack, nil, fmt.Errorf("Не найдено правил для свёртки")
		}
	}
}
```

