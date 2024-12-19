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
