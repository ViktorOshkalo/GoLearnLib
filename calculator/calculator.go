package calculator

func Add[T int | float32 | float64](x, y T) T {
	return x + y
}
