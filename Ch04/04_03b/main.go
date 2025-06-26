package main

import "fmt"

func main() {
	colors := []string{"red", "green", "blue", "yellow", "purple"}

	for i := 0; i < len(colors); i++ {
		println(colors[i])
	}

	for _, v := range colors {
		println(v)
	}

	states := make(map[string]string)
	states["CA"] = "California"
	states["TX"] = "Texas"
	states["NY"] = "New York"

	for k, v := range states {
		println(k, "is", v)
	}

	value := 0
	sum := 0
	for value < 10 {
		sum += value
		value++
		fmt.Println("Current value:", value, "Sum so far:", sum)
	}

	sum = 1
	for sum < 1000 {
		sum *= 2
		fmt.Println("Current sum:", sum)
		if sum > 200 {
			goto theEnd
		}
	}
theEnd:
	fmt.Println("Reached the end of the loop with sum:", sum)
}
