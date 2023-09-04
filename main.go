package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
}

func main() {
	// Register a handler for /hello requests
	http.HandleFunc("/hello", helloHandler)

	// Start the server on port 8080
	fmt.Println("Server is starting on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Server error: %v", err)
	}
}
