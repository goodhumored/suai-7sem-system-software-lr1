package codegenerator

import (
	"fmt"

	"goodhumored/lr1_object_code_generator/code_generator/triad"
	"goodhumored/lr1_object_code_generator/syntax_analyzer/nonterminal"
	"goodhumored/lr1_object_code_generator/syntax_analyzer/parse_tree"
	"goodhumored/lr1_object_code_generator/token"
)

func MapParseTreeToTriadList(tree parse_tree.ParseTree) triad.TriadList {
	triads := triad.NewTriadList()
	_ = mapNodeToTriadList(*tree.Root, &triads)
	return triads
}

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
		fmt.Printf("%v\n", t)
		binaryTriad = &t
	case "xor":
		t := triad.Xor(operand1, operand2, 0)
		binaryTriad = &t
	}
	triads.Add(binaryTriad)
}

func mapAssignment(node parse_tree.Node, triads *triad.TriadList) {
	identifierOperandNode := node.Children[0]
	rightOperandNode := node.Children[2]
	identifierOperand := triad.Id(identifierOperandNode.Value)
	rightOperand := mapNodeToTriadList(*rightOperandNode, triads)
	fmt.Printf("left: %s, right: %s\n", identifierOperand, rightOperand)
	assignmentTriad := triad.Assignment(identifierOperand, rightOperand, 0)
	triads.Add(&assignmentTriad)
}

func mapUnary(node parse_tree.Node, triads *triad.TriadList) {
	operandNode := node.Children[2]
	operand := mapNodeToTriadList(*operandNode, triads)
	notTriad := triad.Not(operand, 0)
	triads.Add(&notTriad)
}
