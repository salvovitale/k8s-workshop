package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	host := os.Getenv("HOSTNAME")
	fmt.Println("Received request")
	fmt.Fprintf(w, "Hello, World! From host: %s", host)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}
