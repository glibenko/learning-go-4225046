package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main Goroutines")
	go say("Hello")
	fmt.Println("main Goroutines")

	go func(message string) {
		fmt.Print(message)
	}("hello from self func")

	time.Sleep(15 * time.Second)
	fmt.Println("main stop")
}

func say(msg string) {
	time.Sleep(5 * time.Second)
	fmt.Println(msg)
}
