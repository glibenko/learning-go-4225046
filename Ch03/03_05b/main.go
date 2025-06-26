package main

import (
	"fmt"
	"sort"
)

func main() {
	states := make(map[string]string)
	states["a"] = "Alabama"
	states["b"] = "Budapest"
	states["c"] = "Costa Rica"
	fmt.Println("Maps", states)

	alabama := states["a"]
	fmt.Println("Alabama:", alabama)
	delete(states, "a")
	fmt.Println("After deleting Alabama:", states)

	states["aa"] = "Denmark"

	keysA := make([]string, 0, len(states))

	for k, v := range states {
		fmt.Printf("%v: %v\n", k, v)
		keysA = append(keysA, k)
	}

	fmt.Println("KeysA:", keysA, len(keysA))

	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println("Sorted keys:", keys)

	for i := range keys {
		fmt.Println(states[keys[i]])
	}
}
