package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	host := os.Getenv("HOSTNAME")
	part := os.Getenv("K8S_WS_PART")
	mode := os.Getenv("K8S_MODE")
	fmt.Println("Received request")
	fmt.Fprintf(w, "Hello, World! From host: %s, part: %s, mode: %s", host, part, mode)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}
