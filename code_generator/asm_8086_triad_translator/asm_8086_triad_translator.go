package asm8086triadtranslator

import (
	"errors"
	"fmt"

	"goodhumored/lr1_object_code_generator/code_generator/triad"
	"goodhumored/lr1_object_code_generator/code_generator/triad/operand"
)

type TriadCodeMap = map[string]string

func JoinCodes(m TriadCodeMap, keys []string) string {
	codes := ""
	for _, key := range keys {
		codes += m[key] + "\n"
	}
	return codes
}

type Asm8086TriadTranslator struct{}

func (t Asm8086TriadTranslator) TranslateTriads(triads []triad.Triad) (string, error) {
	triadCodeMap := TriadCodeMap{}
	mapKeys := []string{}
	for _, triad := range triads {
		mapKeys = append(mapKeys, triad.Hash())
		triadCode, err := translateTriad(triad, triadCodeMap)
		if err != nil {
			return "", err
		}
		triadCodeMap[triad.Hash()] = triadCode
		// fmt.Printf("map: %v\n", triadCodeMap)
	}

	return JoinCodes(triadCodeMap, mapKeys), nil
}

func translateTriad(triadToTranslate triad.Triad, triadCodeMap TriadCodeMap) (string, error) {
	resultCode := ""
	leftLinkOperand, leftOperandIsLink := triadToTranslate.Left().(operand.LinkOperand)
	_, rightOperandIsLink := triadToTranslate.Right().(operand.LinkOperand)
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
			triadLeftOperand, ok := (*leftLinkOperand.LinkTo).(triad.Triad)
			if !ok {
				return "", errors.New("link operand to non triad")
			}
			triadCodeMap[triadLeftOperand.Hash()] += "push ax\n"
			resultCode = fmt.Sprintf("mov dx,ax\npop ax\n%s ax,dx", act)
		} else if leftOperandIsLink {
			resultCode = fmt.Sprintf("%s ax,%s", act, triadToTranslate.Right())
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
	default:
		return "", fmt.Errorf("Неподдерживаемая триада %t", triadToTranslate)
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
