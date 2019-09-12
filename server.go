package main

import (
  "log"
  "net/http"
  "fmt"
  "encoding/json"
)

func main() {
  http.HandleFunc("/time", MakeRequest) // each request calls handler
  log.Fatal(http.ListenAndServe("localhost:8795", nil))
}

func MakeRequest(w http.ResponseWriter, r *http.Request) {
  resp, err := http.Get("http://worldtimeapi.org/api/ip")
  if err != nil {
    log.Fatalln(err)
  }

  var result map[string]interface{}

  json.NewDecoder(resp.Body).Decode(&result)

  fmt.Fprintf(w, "Current time: %q\n", result["datetime"])
}
