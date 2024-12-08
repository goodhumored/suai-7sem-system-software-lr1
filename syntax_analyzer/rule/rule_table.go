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
