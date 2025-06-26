package main

import (
	"fmt"
	"time"
)

func main() {

	weekday := time.Now().Weekday()
	fmt.Printf("Today is %v\n", weekday)

	dayNumber := int(weekday)
	fmt.Printf("The day as a number is %v\n", dayNumber)

	switch dayNumber {
	case 1:
		fmt.Println("It's Monday, the start of the week!")
	case 2:
		fmt.Println("It's Tuesday, the second day of the week!")
	default:
		fmt.Println("It's another day of the week!")
	}
}
