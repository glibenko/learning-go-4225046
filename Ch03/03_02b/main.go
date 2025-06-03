package main

import (
	"fmt"
)

func main() {
	anInt := 42
	var p *int = &anInt
	// p is a pointer to anInt

	if p == nil {
		fmt.Println("p is nil")
	} else {
		fmt.Println("Pointers", *p)
	}

	value1 := 1.3

	pointer1 := &value1
	*pointer1 = *pointer1 + 1.5
	fmt.Println("Pointer to value1:", *pointer1)
	fmt.Println("Value1 after modification:", value1)
}
