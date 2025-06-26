package main

import (
	"fmt"
)

func main() {
	theAnswer := 42
	var result string
	if theAnswer < 0 {
		result = "Less than zero"
	} else if theAnswer == 0 {
		result = "Zero"
	} else {
		result = "Greater than zero"
	}
	fmt.Println("Conditional logic", result)
}
