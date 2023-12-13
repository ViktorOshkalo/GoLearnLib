package calculator

func Add[T int | float32 | float64](x, y T) T {
	return x + y
}

func Distract[T int | float32 | float64](x, y T) T {
	return x - y
}

func Multiply[T int | float32 | float64](x, y T) T {
	return x * y
}

func Divide[T int | float32 | float64](x, y T) T {
	return x / y
}
