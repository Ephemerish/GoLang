// test server
package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	http.HandleFunc("/", hello)

	// Start the server
	go func() {
		fmt.Println("Server listening on http://0.0.0.0:8080")
		if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
			fmt.Printf("Server error: %s\n", err)
		}
	}()

	fmt.Println("Server setup complete. Press Ctrl+C to stop.")
	select {}
}
