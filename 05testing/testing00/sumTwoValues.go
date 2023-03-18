package test00

type Type interface {
	int | float64
}

func SumTwoValues[T Type](x T, y T) T {
	return x + y
}
