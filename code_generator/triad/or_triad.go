package triad

type OrTriad struct {
	LogicTriad
}

func Or(left Operand[int], right Operand[int], number int) OrTriad {
	return OrTriad{
		LogicTriad: Logic(number, left, right, func(left int, right int) int {
			return left | right
		}),
	}
}
