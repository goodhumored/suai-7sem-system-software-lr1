package codegenerator

import "goodhumored/lr1_object_code_generator/code_generator/triad"

// интерфейс переводчика триад
type TriadTranslator interface {
	TranslateTriads(triad.TriadList) (string, error)
}
