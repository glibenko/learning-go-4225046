package main

import (
	"fmt"
)

func main() {
	// This is a slice, not an array
	var colors = []string{"Red", "Green", "Blue"} // fat but not efficient way to create a slice
	var colors2 = make([]string, 0, 3)

	colors2 = append(colors2, "Red", "Creen")
	fmt.Println(colors2)
	colors2 = append(colors2, "Blue", "Black")

	fmt.Println(colors)
	fmt.Println(colors2)

	colors2 = remove(colors2, len(colors2)-1)

	fmt.Println(colors2)
}

func remove(slice []string, i int) []string {
	return append(slice[:i], slice[i+1:]...)
}
