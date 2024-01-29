package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "http://lco.dev"

func main() {
	fmt.Println("LCO web request")
	response, err := http.Get(url)
	checkNilErr(err)

	// fmt.Println("Respose is of type: ", response)
	dataBytes, err := io.ReadAll(response.Body)
	checkNilErr(err)
	content := string(dataBytes)
	fmt.Println(content)

	response.Body.Close() // callers responsiblilty to close the connection
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
