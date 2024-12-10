package asm8086triadtranslator

import (
	"fmt"

	"goodhumored/lr1_object_code_generator/code_generator/triad"
)

type Asm8086TriadTranslator struct{}

func (t Asm8086TriadTranslator) TranslateTriads(triads []triad.Triad) (string, error) {
	code := ""
	for _, triad := range triads {
		triadCode, err := translateTriad(triad)
		if err != nil {
			return code, err
		}
		code += triadCode
	}
	return code, nil
}

func translateTriad(triadToTranslate triad.Triad) (string, error) {
	resultCode := ""
	switch triadToTranslate.(type) {
	case *triad.AssignmentTriad:
		return fmt.Sprintf("mov %s %s\n", triadToTranslate.Left(), triadToTranslate.Right()), nil
	case *triad.AndTriad:
		// leftLinkOperand, leftOperatorIsLink := triadToTranslate.Left().(operand.LinkOperand)
		// rightLinkOperand, rightOperatorIsLink := triadToTranslate.Right().(operand.LinkOperand)
		// if (leftLinkOperand)
		resultCode = fmt.Sprintf("mov ax %s\nand ax %s\n", triadToTranslate.Left(), triadToTranslate.Right())
	case *triad.OrTriad:
		resultCode = fmt.Sprintf("mov ax %s\nor ax %s\n", triadToTranslate.Left(), triadToTranslate.Right())
	case *triad.XorTriad:
		resultCode = fmt.Sprintf("mov ax %s\nxor ax %s\n", triadToTranslate.Left(), triadToTranslate.Right())
	case *triad.NotTriad:
		resultCode = fmt.Sprintf("mov ax %s\nnot ax\n", triadToTranslate.Left())
	default:
		return "", fmt.Errorf("Неподдерживаемая триада %t", triadToTranslate)
	}
	return resultCode, nil
}
