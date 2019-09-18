package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8795", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
  currentTime := time.Now()
  fmt.Fprintf(w, currentTime.String())
}