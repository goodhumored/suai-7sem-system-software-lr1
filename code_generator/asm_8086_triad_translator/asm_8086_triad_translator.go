package asm8086triadtranslator

import (
	"errors"
	"fmt"

	"goodhumored/lr1_object_code_generator/code_generator/triad"
)

type TriadCodeMap = map[int]string

func JoinCodes(m TriadCodeMap) string {
	codes := ""
	for i := 0; i < len(m); i++ {
		codes += m[i]
	}
	return codes
}

type Asm8086TriadTranslator struct{}

func (t Asm8086TriadTranslator) TranslateTriads(triads triad.TriadList) (string, error) {
	triadCodeMap := TriadCodeMap{}
	mapKeys := []string{}
	for _, triad := range triads.Triads() {
		mapKeys = append(mapKeys, triad.Hash())
		triadCode, err := translateTriad(triad, triadCodeMap)
		if err != nil {
			return "", err
		}
		triadCodeMap[triad.Number()] = triadCode
	}

	return JoinCodes(triadCodeMap), nil
}

func translateTriad(triadToTranslate triad.Triad, triadCodeMap TriadCodeMap) (string, error) {
	resultCode := ""
	leftLinkOperand, leftOperandIsLink := triadToTranslate.Left().(triad.LinkOperand)
	_, rightOperandIsLink := triadToTranslate.Right().(triad.LinkOperand)
	switch triadToTranslate.(type) {
	case *triad.AssignmentTriad:
		if rightOperandIsLink {
			return fmt.Sprintf("mov %s,ax\n", triadToTranslate.Left()), nil
		}
		return fmt.Sprintf("mov %s,%s\n", triadToTranslate.Left(), triadToTranslate.Right()), nil
	case *triad.AndTriad, *triad.OrTriad, *triad.XorTriad:
		act, err := getActFromBinaryTriad(triadToTranslate)
		if err != nil {
			return "", err
		}
		if leftOperandIsLink && rightOperandIsLink {
			triadLeftOperand, ok := triadCodeMap[leftLinkOperand.LinkTo]
			if !ok {
				return "", errors.New("link operand to non triad")
			}
			triadLeftOperand += "push ax\n"
			resultCode = fmt.Sprintf("mov dx,ax\npop ax\n%s ax,dx", act)
		} else if leftOperandIsLink {
			resultCode = fmt.Sprintf("%s ax,%s\n", act, triadToTranslate.Right())
		} else if rightOperandIsLink {
			resultCode = fmt.Sprintf("mov dx,ax\nmov ax,%s\n%s ax,dx", triadToTranslate.Left().String(), act)
		} else {
			resultCode = fmt.Sprintf("mov ax,%s\n%s ax,%s\n", triadToTranslate.Left(), act, triadToTranslate.Right())
		}
	case *triad.NotTriad:
		if leftOperandIsLink {
			resultCode += fmt.Sprintf("not ax\n")
		} else {
			resultCode = fmt.Sprintf("mov ax,%s\nnot ax\n", triadToTranslate.Left())
		}
	case *triad.ConstantTriad:
	default:
		return "", fmt.Errorf("Неподдерживаемая триада %v\n", triadToTranslate)
	}
	return resultCode, nil
}

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
