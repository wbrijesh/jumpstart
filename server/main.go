package main

import (
  "fmt"
  "net/http"
  "log"
)

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "hello")
  })
  fmt.Println("Server started on port 4000")
  log.Fatal(http.ListenAndServe(":4000", nil))
}
