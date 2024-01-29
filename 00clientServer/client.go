// test client
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const serverURL = "http://localhost:8080"

func main() {
	// Make a GET request to the server
	response, err := http.Get(serverURL)
	if err != nil {
		handleError("Error making GET request", err)
		return
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		handleError("Error reading response body", err)
		return
	}

	// Print the server response
	fmt.Println("Server Response:", string(body))
}

func handleError(message string, err error) {
	log.Printf("%s: %s\n", message, err)
}
