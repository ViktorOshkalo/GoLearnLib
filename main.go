package main

import (
	"fmt"

	calc "github.com/ViktorOshkalo/GoLearnLib/calculator"
)

func main() {
	// how to do tests?
	add := calc.Add(1, 2.5)
	fmt.Println(add)

	distr := calc.Distract(1, 2.5)
	fmt.Println(distr)

	mult := calc.Multiply(1, 2.5)
	fmt.Println(mult)

	div := calc.Divide(1, 2.5)
	fmt.Println(div)
}
