package triad

type XorTriad struct {
	LogicTriad
}

func Xor(left Operand[int], right Operand[int], number int) XorTriad {
	return XorTriad{
		LogicTriad: Logic(number, left, right, func(left int, right int) int {
			return left ^ right
		}),
	}
}
