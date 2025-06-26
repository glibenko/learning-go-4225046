package main

import (
	"fmt"
)

func main() {
	jack := Dog{"Jack", 12}
	fmt.Println("Structs", jack)
	fmt.Printf("%+v\n", jack)

	jack.Weight = 15
	fmt.Println("Updated Structs", jack)
}

type Dog struct {
	Breed  string
	Weight int
}
