package main

import (
	"fmt"
	"net/http"
	"os"
)

var (
	requestCounter int
	hostname       string
)

func handler(w http.ResponseWriter, r *http.Request) {
	requestCounter++
	fmt.Fprintf(w, "Hi there, got %d requests, running on %s, I love %s!", requestCounter, hostname, r.URL.Path[1:])
}

func main() {
	requestCounter = 0
	hostname = os.Getenv("HOSTNAME")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
