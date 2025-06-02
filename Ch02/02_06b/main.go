package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println("Current time:", t)
	t = time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Date: %s\n", t)

	fmt.Println("Dates and times")

	fmt.Printf(t.Format(time.ANSIC) + "\n")

	fmt.Printf("unix time: %d\n", time.Now().Unix())
	fmt.Printf("iso time"+": %s\n", time.Now().Format(time.RFC3339))

	fmt.Printf("tomorrow: %s\n", time.Now().AddDate(0, 0, 1).Format(time.RFC3339))

	format := "Mon 2006-02-01"
	fmt.Printf("Formatted time: %s\n", t.Format(format))

}
