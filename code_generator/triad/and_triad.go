package triad

type AndTriad struct {
	LogicTriad
}

func And(left Operand[int], right Operand[int], number int) AndTriad {
	return AndTriad{
		LogicTriad: Logic(number, left, right, func(left int, right int) int {
			return left & right
		}),
	}
}
