package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "http://services.explorecalifornia.org/json/tours.php"

func main() {
	fmt.Println("Network requests")
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	checkError(err)

	req.Header.Set("User-Agent", "Go HTTP Client/1.0")

	resp, err := client.Do(req)
	checkError(err)
	defer resp.Body.Close()
	fmt.Println("Response status:", resp.Status)
	fmt.Printf("Response type:%T\n", resp)

	bytes, err := io.ReadAll(resp.Body)
	checkError(err)
	content := string(bytes)
	fmt.Println("Response content:", content)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
