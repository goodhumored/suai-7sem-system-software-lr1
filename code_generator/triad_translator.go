package codegenerator

import "goodhumored/lr1_object_code_generator/code_generator/triad"

type TriadTranslator interface {
	TranslateTriads([]triad.Triad) (string, error)
}
