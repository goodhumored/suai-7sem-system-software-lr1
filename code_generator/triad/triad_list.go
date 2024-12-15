package triad

import "fmt"

type TriadList struct {
	list   []Triad
	length int
}

func (l *TriadList) Add(triad Triad) {
	triad.SetNumber(l.length)
	l.list = append(l.list, triad)
	l.length++
}

func (l TriadList) Triads() []Triad {
	return l.list
}

func (l TriadList) Print() {
	for _, triad := range l.Triads() {
		fmt.Printf("%d)%s\n", triad.Number(), triad.String())
	}
}

func (l TriadList) GetElement(n int) Triad {
	return l.list[n]
}

func (l *TriadList) SetElement(n int, triad Triad) {
	l.list[n] = triad
}

func (l TriadList) Remove(number int) {
	for i := number; i < l.length; i++ {
		triadEl := l.list[i]
		triadEl.SetNumber(triadEl.Number() - 1)
		if operand, ok := triadEl.Left().(LinkOperand); ok {
			operand.LinkTo--
		}
		if operand, ok := triadEl.Right().(LinkOperand); ok {
			operand.LinkTo--
		}
	}
	l.list = append(l.list[:number], l.list[number+1:]...)
	l.length--
}

func (l TriadList) Last() Triad {
	if l.length > 0 {
		return l.list[l.length-1]
	}
	return nil
}

func NewTriadList() TriadList {
	return TriadList{list: []Triad{}, length: 0}
}
