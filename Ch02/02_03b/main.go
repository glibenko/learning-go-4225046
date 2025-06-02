package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	fmt.Print("Enter a text: ")
	str, _ := reader.ReadString('\n')
	fmt.Println(str)
	fmt.Print("Enter another num: ")
	str, _ = reader.ReadString('\n')
	f, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
	if err != nil {
		fmt.Println("Error parsing float:", err)
		return
	} else {
		fmt.Println("Parsed float:", f)
	}
}
