package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Files")
	fileName := "./test.txt"
	file, err := os.Create(fileName)
	defer file.Close()
	checkError(err)
	fmt.Println("File created:", fileName)
	length, err := io.WriteString(file, "Hello, World!\n")
	checkError(err)
	fmt.Println("Wrote", length, "characters to file:", fileName)
	data, err := os.ReadFile(fileName)
	checkError(err)
	fmt.Println("File content:", string(data))
	// content := make([]byte, 64)
	// length, err = file2.Read(content)

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
