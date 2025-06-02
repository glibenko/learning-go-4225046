package main

import (
	"fmt"
)

func main() {
	msg := "Hello from Go!"
	fmt.Println(msg, len(msg))
	msg = "Hello from Go! - modified"
	fmt.Println(msg, len(msg))
	const pi = 3.14

	fmt.Println("Value of pi:", pi)
	var x int = 10
	fmt.Println("Value of x:", x)
	x = 20
	fmt.Println("Modified value of x:", x)
	const y = 30.3
	fmt.Println("Value of y:", y)
}
