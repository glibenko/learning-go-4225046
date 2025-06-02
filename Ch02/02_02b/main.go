package main

import (
	"fmt"
)

func main() {

	str1 := "The quick red fox"
	str2 := "jumped over"
	str3 := "the lazy brown dog."
	num1 := 42

	fmt.Println("Hello from Go!")
	stringLength, err := fmt.Println(str1, str2, str3, num1)
	if err == nil {
		fmt.Println("Total length of printed strings:", stringLength)
	}
	fmt.Printf("Value of num1: %v\n", num1)
	fmt.Printf("Data type of num1: %T\n", num1)
	// fmt.Printf("Data type of num1: %T:\n")
}
