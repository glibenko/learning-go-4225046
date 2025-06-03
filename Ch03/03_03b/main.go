package main

import (
	"fmt"
)

func main() {
	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"
	// colors[3] = "Black"

	fmt.Println("Array", colors)
	fmt.Println("1", colors[1])

	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array", numbers)

	fmt.Println("colors len", len(colors))
	fmt.Println("numbers len", len(numbers))

	fmt.Printf("colors type: %T\n", colors)
}
