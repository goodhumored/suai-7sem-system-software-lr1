package codegenerator

import (
	"fmt"

	"goodhumored/lr1_object_code_generator/code_generator/triad"
	"goodhumored/lr1_object_code_generator/syntax_analyzer/nonterminal"
	"goodhumored/lr1_object_code_generator/syntax_analyzer/parse_tree"
	"goodhumored/lr1_object_code_generator/token"
)

func MapParseTreeToTriadList(tree parse_tree.ParseTree) []triad.Triad {
	triadList, _ := mapNodeToTriadList(*tree.Root)
	return triadList
}

func mapNodeToTriadList(node parse_tree.Node) ([]triad.Triad, triad.Operand) {
	// fmt.Printf("Mapping operand %v\n", node.Symbol.GetName())
	triads := []triad.Triad{}
	var operand triad.Operand
	switch node.Symbol.GetName() {
	case token.IdentifierType.Name:
		// fmt.Printf("its identifier %s\n", node.Value)
		operand = triad.Id(node.Value)
	case nonterminal.Assignment.Name:
		// fmt.Printf("assignment\n")
		triads = append(triads, mapAssignment(node)...)
	case nonterminal.Binary.Name:
		// fmt.Printf("binary\n")
		triads = append(triads, mapBinary(node)...)
	case nonterminal.Unary.Name:
		// fmt.Printf("unary\n")
		triads = append(triads, mapUnary(node)...)
	default:
		for _, child := range node.Children {
			childTriads, childOperand := mapNodeToTriadList(*child)
			triads = append(triads, childTriads...)
			operand = childOperand
		}
	}
	if len(triads) > 0 && operand == nil {
		operand := triad.Link(&triads[len(triads)-1])
		return triads, operand
	}
	return triads, operand
}

func mapBinary(node parse_tree.Node) []triad.Triad {
	triads := []triad.Triad{}
	operator := node.Children[1]
	operandNode1 := node.Children[0]
	operandNode2 := node.Children[2]

	operand1Triads, operand1 := mapNodeToTriadList(*operandNode1)

	triads = append(triads, operand1Triads...)

	operand2Triads, operand2 := mapNodeToTriadList(*operandNode2)

	triads = append(triads, operand2Triads...)

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
	return append(triads, binaryTriad)
}

func mapAssignment(node parse_tree.Node) []triad.Triad {
	identifierOperandNode := node.Children[0]
	rightOperandNode := node.Children[2]
	identifierOperand := triad.Id(identifierOperandNode.Value)
	rightTriads, rightOperand := mapNodeToTriadList(*rightOperandNode)
	fmt.Printf("left: %s, right: %s\n", identifierOperand, rightOperand)
	assignmentTriad := triad.Assignment(identifierOperand, rightOperand, 0)
	return append(rightTriads, &assignmentTriad)
}

func mapUnary(node parse_tree.Node) []triad.Triad {
	operandNode := node.Children[2]
	triads, operand := mapNodeToTriadList(*operandNode)
	notTriad := triad.Not(operand, 0)
	return append(triads, &notTriad)
}
