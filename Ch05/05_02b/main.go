package main

import (
	"fmt"
)

func main() {
	dog := Dog{"Jack", "Woof"}
	dog.Speak()
	dog.Sound = "Bark"
	dog.Speak()

	dog.SpeakMulti(3)
	// fmt.Printf("The %v says %v!\n", dog.Breed, dog.Sound)
}

type Dog struct {
	Breed string
	Sound string
}

func (d Dog) Speak() {
	fmt.Printf("The %v says %v!\n", d.Breed, d.Sound)
}

func (d Dog) SpeakMulti(times int) {
	speach := "The %v says "

	for i := 0; i < times; i++ {
		speach += "%v! "
	}
	fmt.Printf(speach+"\n", d.Breed, d.Sound)
}
