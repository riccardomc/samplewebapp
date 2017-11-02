package main

import (
	"encoding/json"
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

func data(w http.ResponseWriter, r *http.Request) {
	j, err := json.Marshal([]string{"this", "and", "that"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(j)
}

func main() {
	requestCounter = 0
	hostname = os.Getenv("HOSTNAME")
	http.HandleFunc("/", handler)
	http.HandleFunc("/data/", data)
	http.ListenAndServe(":8080", nil)
}
