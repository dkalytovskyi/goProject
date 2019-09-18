package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

//TEST FOR MERGE COMMIT

func main() {
	http.HandleFunc("/time", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8795", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format(time.RFC3339)

	type Message struct {
		Time string `json:"time"`
	}

	b, err := json.Marshal(Message{Time: currentTime})

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Fprintf(w, "%s", b)
}
