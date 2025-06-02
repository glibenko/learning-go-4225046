package main

import "fmt"

func main() {
	i1, i2, i3 := 1, 2, 3
	initSum := i1 + i2 + i3
	fmt.Println("sum:", initSum)

	f1, f2, f3 := 1.3, 1.2, 1.3
	initSumFloat := f1 + f2 + f3
	fmt.Println("sum float:", initSumFloat)

	total := float64(initSum) + initSumFloat
	fmt.Println("total:", total)

	product := float64(initSum) * initSumFloat
	fmt.Println("product:", product)
}
