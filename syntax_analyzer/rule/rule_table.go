package rule

// Таблица правил
type RuleTable struct {
	Rules []Rule
}

// Метод поиска правила по правой части
func (ruleTable RuleTable) GetRuleByRightSide(tokenTypes []Symbol) *Rule {
	for _, rule := range ruleTable.Rules {
		if isApplyable(rule.Right, tokenTypes) {
			return &rule
		}
	}
	return nil
}

// Проверка на применимость правила к целевым символам
func isApplyable(ruleSymbols, targetSymbols []Symbol) bool {
	// Проверяем длины
	lenDiff := len(targetSymbols) - len(ruleSymbols)
	if lenDiff < 0 {
		return false
	}
	// Сравниваем последние символы цепочки символов и символы правила
	for i, ruleSymbol := range ruleSymbols {
		if ruleSymbol.GetName() != targetSymbols[i+lenDiff].GetName() {
			return false
		}
	}
	return true
}
