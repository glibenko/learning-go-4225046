package main

import (
	"fmt"
)

func main() {
	fmt.Println("Functions")
	doSomething() // Calling a private function
}

func doSomething() { //private function
	fmt.Println("Doing something")
	sum := addValues(5, 10)
	fmt.Println("Sum:", sum)
	sumAll := addAll(1, 2, 3, 4, 5)
	fmt.Println("Sum of all values:", sumAll)
	sumAll2, _, _ := addAll2(1, 2, 3, 4, 5)
	fmt.Println("Sum of all values:", sumAll2)
}

func addValues(val1, val2 int) int {
	return val1 + val2
}

func DoingSomethingElse() { // Exported function
	fmt.Println("Doing something else")
}

func addAll(values ...int) int {
	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum
}

func addAll2(values ...int) (int, int, float64) {
	sum := 0
	for _, value := range values {
		sum += value
	}
	counst := len(values)
	average := float64(sum) / float64(counst)
	return sum, counst, average
}
